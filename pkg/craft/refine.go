package craft

import "fmt"

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
		fmt.Println("Inventário sem material de entrada")
		delete(m.inventory.Materials, d.Input)
		return
	}

	fmt.Printf("Inventário com %f de %s\n", m.inventory.Materials[d.Input], material.Name)

	substances := material.Reduce(1)

	for id, quantity := range substances {
		substance, _ := GetSubstance(id)
		elements := substance.Reduce(quantity)
		for id, quantity := range elements {
			element, _ := GetElement(id)
			if m.Heat >= element.BoilingPoint {
				deltaTemp := m.Heat - element.BoilingPoint
				currentLoss := clamp(deltaTemp*element.Volatility*dt, 0.0, 1.0)
				m.AddElementToInventory(id, quantity*(1-currentLoss))
			} else {
				m.AddElementToInventory(id, quantity)
			}
		}
	}
}

func clamp(v, min, max float32) float32 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
