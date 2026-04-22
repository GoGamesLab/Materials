package craft

import (
	"fmt"
	"math"
)

// Basicamente refinar é processar um material e decompor ele em seus constituintes

type RefineOperation struct {
	Operation
	Input  MaterialID  // Aqui entram os produtos necessários
	Output []Composite // Aqui saem as substâncias desejadas E os resíduos/excedentes (Escória/Gases)
}

func (d RefineOperation) GetOperation() Operation { return d.Operation }
func (d RefineOperation) Kind() string            { return "refine" }

func (m *Machine) finishRefination(d RefineOperation, dt float32) {
	material, _ := GetMaterial(d.Input)
	m.inventory.Materials[d.Input] -= 1
	if m.inventory.Materials[d.Input] <= 0 {
		delete(m.inventory.Materials, d.Input)
		return
	}

	substances := material.Reduce(1)

	for id, quantity := range substances {
		substance, _ := GetSubstance(id)
		fmt.Printf("refinando substância %s\n", substance.Name)
		elements := substance.Reduce(quantity)
		for id, quantity := range elements {
			element, _ := GetElement(id)
			fmt.Printf("separando elemento %s\n", element.Name)
			if m.Heat >= element.BoilingPoint {
				currentLoss := (1.0 - float32(math.Exp(-0.5*float64(m.Heat-element.BoilingPoint)))) * element.Volatility * 100.0 * dt
				// aqui, adicionar a quantidade se já existir no inventário
				m.ConsumeElement(id, currentLoss)
			} else {
				// não esquentou o suficiente, acumula no inventário
				// mas quando começar a esquentar deve queimar o inventário também!
				m.ProduceElement(id, quantity)
			}
		}
	}
}
