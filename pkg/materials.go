package materials

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

// const (
// 	AirID MaterialID = iota
// 	CoalID
// 	CrudeOilID
// 	IronBarID
// 	IronPlateID
// 	SteelID
// 	UraniumOreID
// 	YellowcakeID
// 	LEUFuelRodID
// 	DepletedUraniumRodID
// )

var (
	Materials  = make(map[MaterialID]Material)
	Signatures = make(map[string][]MaterialID)
)

func init() {
	loadMaterialsFromJSON("../../Grind/assets/materials/materials.json")
	// RegisterMaterial(Material{
	// 	ID:   AirID,
	// 	Name: "Air",
	// 	composites: []Composite{
	// 		{PureNitrogenID, 78},
	// 		{PureOxygenID, 21},
	// 		{PureArgonID, 1},
	// 		{CarbonDioxideID, 0.5},
	// 	},
	// })
	// RegisterMaterial(Material{
	// 	ID:   CoalID,
	// 	Name: "Coal",
	// 	composites: []Composite{
	// 		{PureCarbonID, 80},
	// 		{PureSulfurID, 5},
	// 		{MethaneID, 15},
	// 	},
	// 	State:       Solid,
	// 	Temperature: 21,  // Temp ambiente inicial
	// 	Hardness:    2.5, // Fácil de minerar
	// })
	// RegisterMaterial(Material{
	// 	ID:   CrudeOilID,
	// 	Name: "Crude Oil",
	// 	composites: []Composite{
	// 		{ParaffinID, 60},   // Fração pesada (Betume/Ceras)
	// 		{MethaneID, 30},    // Fração leve (Gases dissolvidos)
	// 		{PureSulfurID, 10}, // Impurezas ácidas (Enxofre)
	// 	},
	// 	State:       Liquid,
	// 	Temperature: 21,  // Temperatura ambiente
	// 	Hardness:    0.0, // Não resiste à mineração (é bombeado)
	// })
	// RegisterMaterial(Material{
	// 	ID:   IronBarID,
	// 	Name: "Barra de ferro",
	// 	composites: []Composite{
	// 		{PureIronID, 98},  // Majoritariamente Ferro
	// 		{PureSulfurID, 1}, // Com um pouco de impurezas
	// 		{PureSiliconID, 0.5},
	// 		{PureManganeseID, 0.25},
	// 		{PurePhosphorusID, 0.25},
	// 	},
	// 	State:       Solid,
	// 	Temperature: 21,
	// 	Hardness:    7.0, // Exige ferramentas avançadas
	// })
	// RegisterMaterial(Material{
	// 	ID:   IronPlateID,
	// 	Name: "Chapa de ferro",
	// 	composites: []Composite{
	// 		{PureIronID, 98},  // Majoritariamente Ferro
	// 		{PureSulfurID, 1}, // Com um pouco de impurezas
	// 		{PureSiliconID, 0.5},
	// 		{PureManganeseID, 0.25},
	// 		{PurePhosphorusID, 0.25},
	// 	},
	// 	State:       Solid,
	// 	Temperature: 21,
	// 	Hardness:    7.0, // Exige ferramentas avançadas
	// })
	// RegisterMaterial(Material{
	// 	ID:   SteelID,
	// 	Name: "Steel",
	// 	composites: []Composite{
	// 		{PureIronID, 98},  // O aço é majoritariamente Ferro
	// 		{PureCarbonID, 2}, // Com um pouco de Carbono para dureza
	// 	},
	// 	State:       Solid,
	// 	Temperature: 21,
	// 	Hardness:    7.0, // Exige ferramentas avançadas
	// })
	// RegisterMaterial(Material{
	// 	ID: UraniumOreID, Name: "Uranium Ore",
	// 	composites: []Composite{
	// 		{PureSiliconID, 95},   // Pedra/Quartzo ao redor
	// 		{UraniumDioxideID, 5}, // Apenas 5% de Urânio real
	// 	},
	// 	State:    Solid,
	// 	Hardness: 4.5,
	// })
	// // Yellowcake (Urânio processado e concentrado obtido na refinaria química)
	// RegisterMaterial(Material{
	// 	ID: YellowcakeID, Name: "Yellowcake",
	// 	composites: []Composite{
	// 		{UraniumDioxideID, 100},
	// 	},
	// 	State:    Solid,
	// 	Hardness: 1.5,
	// })
	// // Célula de Combustível Nuclear Enriquecido (Pronto para o Reator)
	// RegisterMaterial(Material{
	// 	ID: LEUFuelRodID, Name: "Enriched Uranium Fuel Rod",
	// 	composites: []Composite{
	// 		{UraniumDioxideID, 90}, // Urânio denso enriquecido simulado
	// 		{PureIronID, 10},       // Revestimento protetor metálico (Zircaloy simplificado para Iron)
	// 	},
	// 	State:    Solid,
	// 	Hardness: 6.0,
	// })
	// // Barra de Combustível Exaurida (O Lixo Nuclear)
	// RegisterMaterial(Material{
	// 	ID: DepletedUraniumRodID, Name: "Depleted Uranium Fuel Rod",
	// 	composites: []Composite{
	// 		{Cesium137ID, 60},      // Subproduto de fissão perigoso
	// 		{PlutoniumOxideID, 10}, // Transmutação de U-238 em Plutônio
	// 		{PureIronID, 30},
	// 	},
	// 	State:    Solid,
	// 	Hardness: 6.0,
	// })
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
		fmt.Fprintf(&sb, "%s|", c.Substance)
	}
	return sb.String()
}

func RegisterMaterial(m Material) error {
	if _, exists := Materials[m.ID]; exists {
		return fmt.Errorf("🧨 Material com ID %s já registrada", m.ID)
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

func (m *Material) Reduction(quantity float32) map[SubstanceID]float32 {
	r := make(map[SubstanceID]float32)
	for _, c := range m.Composites {
		r[c.Substance] = (c.Percentual / 100) * quantity
	}

	return r
}

func Compositions(composite []Composite) ([]Material, error) {
	sig := GenerateSignature(composite)
	ids, ok := Signatures[sig]
	if !ok || len(ids) == 0 {
		return nil, fmt.Errorf("🧨 No material found for signature: %s", sig)
	}
	var result []Material
	for _, id := range ids {
		result = append(result, Materials[id])
	}

	return result, nil
}

func loadMaterialsFromJSON(materialsPath string) error {

	// Carregar Materiais
	mData, err := os.ReadFile(materialsPath)
	if err != nil {
		return fmt.Errorf("erro lendo materiais: %w", err)
	}
	var mList []Material
	if err := json.Unmarshal(mData, &mList); err != nil {
		return err
	}
	for _, m := range mList {
		Materials[m.ID] = m
		sig := GenerateSignature(m.Composites)
		Signatures[sig] = append(Signatures[sig], m.ID)
	}

	return nil
}
