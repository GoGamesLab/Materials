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

func registerOp[T Procedure](store map[OperationID]T, op T) error {
	id := op.GetOperation().ID
	if _, exists := store[id]; exists {
		return fmt.Errorf("🧨 Operação com ID %d já registrada", id)
	}
	store[id] = op
	return nil
}

func RegisterOperation(p Procedure) error {
	machinesMutex.Lock()
	defer machinesMutex.Unlock()

	switch op := p.(type) {
	case RefineOperation:
		return registerOp(RefineOperations, op)
	case SynthesizeOperation:
		return registerOp(SynthesizeOperations, op)
	case TransformOperation:
		return registerOp(TransformOperations, op)
	default:
		return fmt.Errorf("🧨 Tipo de operação desconhecida: %T", p)
	}
}

func GetMachine(id MachineID) Machine {
	if m, ok := Machines[id]; ok {
		return m
	}
	return Machines[EntropyID] // Fallback seguro
}
