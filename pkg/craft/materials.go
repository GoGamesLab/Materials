package craft

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type PhysicalState uint8

const (
	Solid  PhysicalState = 1
	Liquid PhysicalState = 2
	Gas    PhysicalState = 3
	Plasma PhysicalState = 4
)

type Composite struct {
	substance  SubstanceID
	percentual float32
}

type Material struct {
	ID            MaterialID // 2 bytes - Aponta para MaterialDefinition
	Name          string
	Composites    []Composite
	State         PhysicalState // 1 byte  - Estado atual do bloco
	HP            uint8         // 1 byte  - Integridade (0-255)
	Temperature   float32       // 2 bytes - Temperatura local para simular mudanças de estado
	Hardness      float32       // Para mecânica de mineração
	Density       float32
	SpecificHeat  float32 // Calor específico: quanto de energia precisa para mudar 1°C
	EnergyDensity float32 // "Calorias": Joules por unidade de massa se queimado/reagido
	FlashPoint    float32 // Temperatura em que o material entra em combustão
	Conductivity  float32 // Quão rápido ele transfere calor para vizinhos
}

const (
	AirID      MaterialID = 0
	CoalID     MaterialID = 1
	CrudeOilID MaterialID = 2
	SteelID    MaterialID = 3
)

var (
	Materials      = make(map[MaterialID]Material)
	Signatures     = make(map[string][]MaterialID)
	materialsMutex sync.Mutex
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
}

func GenerateSignature(composites []Composite) string {
	// 1. Criar uma cópia para não mexer no original e ordenar por ID
	// Isso garante que a assinatura seja determinística
	temp := make([]Composite, len(composites))
	copy(temp, composites)

	sort.Slice(temp, func(i, j int) bool {
		return temp[i].substance < temp[j].substance
	})

	// 2. Montar a string: "SubstanceID:Percentage|..."
	var sb strings.Builder
	for _, c := range temp {
		fmt.Fprintf(&sb, "%d|", c.substance)
	}
	return sb.String()
}

func RegisterMaterial(m Material) error {
	materialsMutex.Lock()
	defer materialsMutex.Unlock()

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
		r[c.substance] = (c.percentual / 100) * quantity
	}

	return r
}
