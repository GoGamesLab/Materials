package materials

import (
	"fmt"
	"sort"
	"strings"
)

const (
	Solid SubstanceState = iota + 1
	Liquid
	Gas
	Plasma
)

const (
	AirID MaterialID = iota
	CoalID
	CrudeOilID
	IronBarID
	IronPlateID
	SteelID
	UraniumOreID
	YellowcakeID
	LEUFuelRodID
	DepletedUraniumRodID
)

var (
	Materials  = make(map[MaterialID]Material)
	Signatures = make(map[string][]MaterialID)
)

func init() {
	RegisterMaterial(Material{
		ID:   AirID,
		Name: "Air",
		Composites: []Composite{
			{PureNitrogenID, 78},
			{PureOxygenID, 21},
			{PureArgonID, 1},
			{CarbonDioxideID, 0.5},
		},
	})
	RegisterMaterial(Material{
		ID:   CoalID,
		Name: "Coal",
		Composites: []Composite{
			{PureCarbonID, 80},
			{PureSulfurID, 5},
			{MethaneID, 15},
		},
		State:         Solid,
		HP:            80,    // Carvão é quebradiço
		Temperature:   21,    // Temp ambiente inicial
		Hardness:      2.5,   // Fácil de minerar
		Density:       1.5,   // Leve
		SpecificHeat:  0.9,   // Esquenta relativamente rápido
		EnergyDensity: 24.0,  // Valor alto para combustão (MJ/kg)
		FlashPoint:    300.0, // Acende com calor moderado
		Conductivity:  0.2,   // Pobre condutor (ajuda a manter o calor na fornalha)
	})
	RegisterMaterial(Material{
		ID:   CrudeOilID,
		Name: "Crude Oil",
		Composites: []Composite{
			{ParaffinID, 60},   // Fração pesada (Betume/Ceras)
			{MethaneID, 30},    // Fração leve (Gases dissolvidos)
			{PureSulfurID, 10}, // Impurezas ácidas (Enxofre)
		},
		State:         Liquid,
		HP:            100,  // Líquidos não "quebram" como sólidos, mas podem ser "dispersos"
		Temperature:   21,   // Temperatura ambiente
		Hardness:      0.0,  // Não resiste à mineração (é bombeado)
		Density:       0.9,  // Menos denso que a água (boia)
		SpecificHeat:  2.0,  // Alta inércia térmica (demora para esquentar/esfriar)
		EnergyDensity: 45.0, // Muito alta (quase o dobro do carvão por massa)
		FlashPoint:    60.0, // Perigoso: entra em combustão a temperaturas baixas
		Conductivity:  0.1,  // Isolante térmico (comum em óleos)
	})
	RegisterMaterial(Material{
		ID:   SteelID,
		Name: "Steel",
		Composites: []Composite{
			{PureIronID, 98},  // O aço é majoritariamente Ferro
			{PureCarbonID, 2}, // Com um pouco de Carbono para dureza
		},
		State:         Solid,
		HP:            255, // Máxima durabilidade
		Temperature:   21,
		Hardness:      7.0,    // Exige ferramentas avançadas
		Density:       7.8,    // Muito pesado
		SpecificHeat:  0.5,    // Esquenta muito rápido (bom para caldeiras)
		EnergyDensity: 0.0,    // Aço não queima (geralmente)
		FlashPoint:    1400.0, // Ponto de amolecimento/fusão
		Conductivity:  0.8,    // Ótimo condutor para transferir calor para a água
	})
	RegisterMaterial(Material{
		ID: UraniumOreID, Name: "Uranium Ore",
		Composites: []Composite{
			{Substance: PureSiliconID, Percentual: 95},   // Pedra/Quartzo ao redor
			{Substance: UraniumDioxideID, Percentual: 5}, // Apenas 5% de Urânio real
		},
		State: Solid, HP: 120, Hardness: 4.5, Density: 3.5,
		SpecificHeat: 0.8, EnergyDensity: 0, FlashPoint: 0, Conductivity: 0.2,
	})

	// Yellowcake (Urânio processado e concentrado obtido na refinaria química)
	RegisterMaterial(Material{
		ID: YellowcakeID, Name: "Yellowcake",
		Composites: []Composite{
			{Substance: UraniumDioxideID, Percentual: 100},
		},
		State: Solid, HP: 50, Hardness: 1.5, Density: 5.5,
		SpecificHeat: 0.4, EnergyDensity: 0, FlashPoint: 0, Conductivity: 0.1,
	})

	// Célula de Combustível Nuclear Enriquecido (Pronto para o Reator)
	RegisterMaterial(Material{
		ID: LEUFuelRodID, Name: "Enriched Uranium Fuel Rod",
		Composites: []Composite{
			{Substance: UraniumDioxideID, Percentual: 90}, // Urânio denso enriquecido simulado
			{Substance: PureIronID, Percentual: 10},       // Revestimento protetor metálico (Zircaloy simplificado para Iron)
		},
		State: Solid, HP: 200, Hardness: 6.0, Density: 19.1,
		SpecificHeat:  0.12,      // Esquenta incrivelmente rápido
		EnergyDensity: 3900000.0, // Densidade energética absurda em MJ/kg comparado ao carvão (24.0)
		FlashPoint:    2500.0,
		Conductivity:  0.8,
	})

	// Barra de Combustível Exaurida (O Lixo Nuclear)
	RegisterMaterial(Material{
		ID: DepletedUraniumRodID, Name: "Depleted Uranium Fuel Rod",
		Composites: []Composite{
			{Substance: Cesium137ID, Percentual: 60},      // Subproduto de fissão perigoso
			{Substance: PlutoniumOxideID, Percentual: 10}, // Transmutação de U-238 em Plutônio
			{Substance: PureIronID, Percentual: 30},
		},
		State: Solid, HP: 200, Hardness: 6.0, Density: 18.0,
		SpecificHeat: 0.15, EnergyDensity: 0, FlashPoint: 0, Conductivity: 0.5,
	})

}

func GenerateSignature(composites []Composite) string {
	// 1. Criar uma cópia para não mexer no original e ordenar por ID
	// Isso garante que a assinatura seja determinística
	temp := make([]Composite, len(composites))
	copy(temp, composites)

	sort.Slice(temp, func(i, j int) bool {
		return temp[i].Substance < temp[j].Substance
	})

	// 2. Montar a string: "SubstanceID:Percentage|..."
	var sb strings.Builder
	for _, c := range temp {
		fmt.Fprintf(&sb, "%d|", c.Substance)
	}
	return sb.String()
}

func RegisterMaterial(m Material) error {
	if _, exists := Materials[m.ID]; exists {
		return fmt.Errorf("🧨 Material com ID %d já registrada", m.ID)
	}

	// 1. Gera a assinatura antes de salvar
	sig := GenerateSignature(m.Composites)

	// 2. Registra nos dois mapas
	Materials[m.ID] = m
	Signatures[sig] = append(Signatures[sig], m.ID)

	return nil
}

func GetMaterial(id MaterialID) (*Material, error) {
	if m, ok := Materials[id]; ok {
		return &m, nil
	}
	return nil, fmt.Errorf("🧨 Material %v: not found", id)
}

func (m *Material) Reduce(quantity float32) map[SubstanceID]float32 {
	r := make(map[SubstanceID]float32)
	for _, c := range m.Composites {
		r[c.Substance] = (c.Percentual / 100) * quantity
	}

	return r
}
