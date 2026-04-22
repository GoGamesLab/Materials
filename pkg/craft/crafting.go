package craft

import (
	"math"
)

// ID único para busca rápida
type ElementID uint16
type SubstanceID uint16
type MaterialID uint16
type OperationID uint16
type ProcedureID uint16
type MachineID uint16
type ProductID uint16

type Operation struct {
	ID               OperationID
	Name             string
	RequiredTemp     float32 // Condição térmica para manter a reação
	EnergyCost       float32 // Energia (Joules) necessária para a reação
	Duration         float32 // Tempo em segundos
	BaseEnergyChange float32 // Negativo = Consome | Positivo = Produz
	ActivationTemp   float32 // Temperatura mínima para a reação começar
	Dissipation      float32 // Índice de dissipação
}

type MachineType int

const (
	Manual MachineType = iota
	Mixer
	Synthesizer
	Processor
	Extruder
	Laminator
	Refiner
	Assembler
	Smelter
)

type Procedure interface {
	GetOperation() Operation
	Kind() string
}

type ProductType int

const (
	MaterialType ProductType = iota
	SubstanceType
)

type Inventory struct {
	Materials  map[MaterialID]float32
	Substances map[SubstanceID]float32
	Elements   map[ElementID]float32
}

type Machine struct {
	ID         MachineID
	Name       string
	Heat       float32 // Temperatura interna atual
	Procedures []Procedure
	Progress   float32
	inventory  Inventory
}

func (m *Machine) Update(dt float32) {
	for _, step := range m.Procedures {
		p := step.GetOperation()

		// Check de Ativação
		if m.Heat >= p.ActivationTemp {
			// Execução
			m.executeStep(step, dt)

			m.Heat = calculateTemperature(m.Heat, p.BaseEnergyChange*dt, p.Dissipation)
		} else {
			// Balanço Térmico
			m.Heat += p.BaseEnergyChange * dt
		}
	}
}

func (m *Machine) GetInventory() Inventory {
	return Inventory{
		Materials:  m.inventory.Materials,
		Substances: m.inventory.Substances,
		Elements:   m.inventory.Elements,
	}
}

func (m *Machine) executeStep(step Procedure, dt float32) {
	p := step.GetOperation()

	// 1. Acumula o progresso baseado no tempo real que passou
	m.Progress += dt

	// 2. Verifica se o tempo de duração foi atingido
	// isso aqui não está sendo um processo contínuo
	// mas tem um desempenho melhor (só calcular tudo ao final do tempo)
	if m.Progress >= p.Duration {
		// O PROCESSO TERMINOU: Hora de realizar a operação
		switch v := step.(type) {
		case RefineOperation:
			m.finishRefination(v, dt)
		case SynthesizeOperation:
			m.finishSynthesization(v, dt)
		case TransformOperation:
			m.finishTransformation(v, dt)
		}

		// 3. Reseta o progresso para o próximo ciclo (ou próximo passo da chain)
		m.Progress = 0
	}
}

// calculateTemperature calcula a temperatura no próximo tick.
// current: temperatura atual
// heatGain: ganho de temperatura por tick enquanto acesa
// retorna: nova temperatura após um tick
func calculateTemperature(current, heatGain float32, dissipation float32) float32 {
	const Tenv = 21.0

	// Tempo de meia-vida em ticks para a diferença T - Tenv.
	// Ajuste aqui para alterar quão rápido a fornalha dissipa calor.
	// Ex.: dissipation = 50 → dissipa mais devagar; dissipation = 10 → mais rápido.

	// taxa de dissipação por tick derivada da meia-vida:
	// fator de retenção por tick = 2^(-1/tHalf)
	retain := math.Pow(2.0, -1.0/float64(dissipation))
	k := 1.0 - retain // fração perdida por tick

	// Atualiza: acumula ganho e aplica dissipação em relação ao ambiente
	delta := float64(current - Tenv)
	delta = delta*(1.0-k) + float64(heatGain) // primeiro aplica retenção, depois adiciona ganho
	// equivalência ao modelo T_next = Tenv + (T_current - Tenv)*(1-k) + heatGain

	return float32(Tenv + delta)
}

type ResourceID interface {
	ElementID | SubstanceID | MaterialID
}

func getInventoryMap[K ResourceID](m *Machine) *map[K]float32 {
	var anyMap any
	switch any(new(K)).(type) {
	case *ElementID:
		anyMap = &m.inventory.Elements
	case *SubstanceID:
		anyMap = &m.inventory.Substances
	case *MaterialID:
		anyMap = &m.inventory.Materials
	}
	return anyMap.(*map[K]float32)
}

func Store[K ResourceID](m *Machine, id K, amount float32) {
	if amount <= 0 {
		return
	}

	invMap := getInventoryMap[K](m)
	if *invMap == nil {
		*invMap = make(map[K]float32)
	}

	unit := amount
	(*invMap)[id] += unit
}

func Produce[K ResourceID](m *Machine, id K, amount float32) {
	if amount <= 0 {
		return
	}

	invMap := getInventoryMap[K](m)
	if *invMap == nil {
		*invMap = make(map[K]float32)
	}

	unit := amount / 100
	(*invMap)[id] += unit
}

func Consume[K ResourceID](m *Machine, id K, amount float32) {
	if amount <= 0 {
		return
	}

	invMap := getInventoryMap[K](m)
	if *invMap == nil {
		return // Ou m.Produce(id, amount) se quiser manter o comportamento original de 'negativar'
	}

	unit := amount / 100
	(*invMap)[id] -= unit

	if (*invMap)[id] <= 0 {
		delete(*invMap, id)
	}
}
