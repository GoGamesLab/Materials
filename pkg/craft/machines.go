package craft

import (
	"fmt"
	"sync"
)

// Exemplo de ID para o processo
const (
	CoalCombustionID     OperationID = 100
	CrudeOilCombustionID OperationID = 101
)

const (
	EntropyID        MachineID = 0
	CoalBurnerID     MachineID = 1
	CrudeOilBurnerID MachineID = 2
)

var (
	Machines             = make(map[MachineID]Machine)
	RefineOperations     = make(map[OperationID]RefineOperation)
	SynthesizeOperations = make(map[OperationID]SynthesizeOperation)
	TransformOperations  = make(map[OperationID]TransformOperation)
	machinesMutex        sync.RWMutex
)

func init() {
	RegisterOperation(RefineOperation{
		Operation: Operation{
			ID:               CoalCombustionID,
			Name:             "Combustão de Carvão",
			RequiredTemp:     600.0, // Temperatura ideal de queima
			ActivationTemp:   300.0, // FlashPoint do seu Carvão
			Duration:         10.0,  // Cada unidade de carvão dura 10 segundos
			EnergyCost:       0,     // Não gasta energia elétrica/externa
			BaseEnergyChange: 150.0, // Produz calor para a máquina e arredores
			Dissipation:      3.5,   // Perda para o ambiente (0 = 100%)
		},
		// O que a fogueira "consome" do inventário
		Input: CoalID,
		// O que sobra (Cinzas/Carbono e Gases para a atmosfera) deve ser jogado
		// inventário da máquina
		Output: []Composite{
			{PureCarbonID, 5}, // Sobra um pouco de cinza/resíduo
			{PureSulfurID, 5}, // Libera enxofre (poluição)
			{MethaneID, 15},   // Libera os voláteis
		},
	})
	RegisterOperation(RefineOperation{
		Operation: Operation{
			ID:               CrudeOilCombustionID,
			Name:             "Combustão de Óleo Crú",
			RequiredTemp:     700.0, // Temperatura ideal de queima
			ActivationTemp:   400.0, // FlashPoint do seu Óleo
			Duration:         15.0,  // Cada unidade de óleo dura 5 segundos
			EnergyCost:       0,     // Não gasta energia elétrica/externa
			BaseEnergyChange: 250.0, // Produz calor para a máquina e arredores
			Dissipation:      3.5,   // Perda para o ambiente (0 = 100%)
		},
		// O que a fogueira "consome" do mundo
		Input: CrudeOilID,
		// O que sobra (Cinzas/Carbono e Gases para a atmosfera) deve ser jogado
		// inventário da máquina
		Output: []Composite{
			{ParaffinID, 2},
			{MethaneID, 5},
			{PureSulfurID, 3},
		},
	})

	RegisterMachine(Machine{
		ID:         EntropyID,
		Name:       "Entropia",
		Heat:       21.0,
		Procedures: []Procedure{},
	})
	RegisterMachine(Machine{
		ID:   CoalBurnerID,
		Name: "Fogueira de Acampamento",
		Heat: 310.0, // Iniciada com um fósforo (acima dos 300°C de ativação)
		Procedures: []Procedure{
			RefineOperations[CoalCombustionID], // A fogueira está configurada para queimar carvão
		},
	})
	RegisterMachine(Machine{
		ID:   CrudeOilBurnerID,
		Name: "Tambor de Óleo",
		Heat: 310.0, // Iniciada com um fósforo (abaixo dos 400°C de ativação)
		Procedures: []Procedure{
			RefineOperations[CrudeOilCombustionID], // A fogueira está configurada para queimar óleo crú
		},
	})
}

func RegisterMachine(m Machine) error {
	machinesMutex.Lock()
	defer machinesMutex.Unlock()

	if _, exists := Machines[m.ID]; exists {
		return fmt.Errorf("🧨 Máquina com ID %d já registrada", m.ID)
	}

	Machines[m.ID] = m

	return nil
}

func RegisterOperation(p Procedure) error {
	machinesMutex.Lock()
	defer machinesMutex.Unlock()

	id := p.GetOperation().ID
	switch p := p.(type) {
	case RefineOperation:
		if _, exists := RefineOperations[id]; exists {
			return fmt.Errorf("🧨 Operação de refinar com ID %d já registrada", id)
		}

		RefineOperations[id] = p
	case SynthesizeOperation:
		if _, exists := SynthesizeOperations[id]; exists {
			return fmt.Errorf("🧨 Operação de sintetizar com ID %d já registrada", id)
		}

		SynthesizeOperations[id] = p
	case TransformOperation:
		if _, exists := TransformOperations[id]; exists {
			return fmt.Errorf("🧨 Operação de transformar com ID %d já registrada", id)
		}

		TransformOperations[id] = p
	default:
		return fmt.Errorf("🧨 Tipo de operação desconhecida: %T", p)
	}

	return nil
}

func GetMachine(id MachineID) Machine {
	if m, ok := Machines[id]; ok {
		return m
	}
	return Machines[EntropyID] // Fallback seguro
}

func (m *Machine) AddElementToInventory(id ElementID, amount float32) {
	// A produção considera percentual
	m.ProduceElement(id, amount*100)
}

func (m *Machine) AddSubstanceToInventory(id SubstanceID, amount float32) {
	// A produção considera percentual
	m.ProduceSubstance(id, amount*100)
}

func (m *Machine) AddMaterialToInventory(id MaterialID, amount float32) {
	// A produção considera percentual
	m.ProduceMaterial(id, amount*100)
}

func (m *Machine) ProduceElement(id ElementID, amount float32) {
	if amount <= 0 {
		return
	}
	if m.inventory.Elements == nil {
		m.inventory.Elements = make(map[ElementID]float32)
	}
	unit := 1 * amount / 100
	m.inventory.Elements[id] += unit
}

func (m *Machine) ProduceSubstance(id SubstanceID, amount float32) {
	if amount <= 0 {
		return
	}
	if m.inventory.Substances == nil {
		m.inventory.Substances = make(map[SubstanceID]float32)
	}
	unit := 1 * amount / 100
	m.inventory.Substances[id] += unit
}

func (m *Machine) ProduceMaterial(id MaterialID, amount float32) {
	if amount <= 0 {
		return
	}
	if m.inventory.Materials == nil {
		m.inventory.Materials = make(map[MaterialID]float32)
	}
	unit := 1 * amount / 100
	m.inventory.Materials[id] += unit
}

func (m *Machine) ConsumeElement(id ElementID, amount float32) {
	if amount <= 0 {
		return
	}
	if m.inventory.Elements == nil {
		m.ProduceElement(id, amount)
	}
	unit := 1 * amount / 100
	m.inventory.Elements[id] -= unit
	if m.inventory.Elements[id] <= 0 {
		delete(m.inventory.Elements, id)
	}
}

func (m *Machine) ConsumeSubstance(id SubstanceID, amount float32) {
	if amount <= 0 {
		return
	}
	if m.inventory.Substances == nil {
		m.ProduceSubstance(id, amount)
	}
	unit := 1 * amount / 100
	m.inventory.Substances[id] -= unit
	if m.inventory.Substances[id] <= 0 {
		delete(m.inventory.Substances, id)
	}
}

func (m *Machine) ConsumeMaterial(id MaterialID, amount float32) {
	if amount <= 0 {
		return
	}
	if m.inventory.Materials == nil {
		m.ProduceMaterial(id, amount)
	}
	unit := 1 * amount / 100
	m.inventory.Materials[id] -= unit
	if m.inventory.Materials[id] <= 0 {
		delete(m.inventory.Materials, id)
	}
}
