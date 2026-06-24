package materials

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Substances = make(map[SubstanceID]*Substance)
)

func RegisterSubstance(s Substance) error {
	if _, exists := Substances[s.ID]; exists {
		return fmt.Errorf("🧨 Substância com ID %s já registrada", s.ID)
	}
	Substances[s.ID] = &s

	return nil
}

func GetSubstance(id SubstanceID) (*Substance, error) {
	if s, ok := Substances[id]; ok {
		return s, nil
	}

	return nil, fmt.Errorf("🧨 Substance %v: not found", id)
}

func (s Substance) GetMolecularWeight() float64 {
	var total float64
	for _, bond := range s.Composition {
		element := Elements[bond.Element]
		total += element.Weight * float64(bond.Amount)
	}
	return total
}

type State int

func (s Substance) GetState(currentTemp float32) SubstanceState {
	if currentTemp < s.MeltingPoint {
		return Solid
	}
	if currentTemp < s.BoilingPoint {
		return Liquid
	}
	return Gas
}

func (s *Substance) Reduce(quantity float32) map[ElementID]float32 {
	r := make(map[ElementID]float32)
	for _, c := range s.Composition {
		r[c.Element] = float32(c.Amount) * quantity
	}

	return r
}

func LoadSubstancesFromJSON(substancesPath string) error {
	sData, err := os.ReadFile(substancesPath)
	if err != nil {
		return err
	}

	var sList []Substance
	if err := json.Unmarshal(sData, &sList); err != nil {
		return err
	}

	for _, s := range sList {
		subs := s
		if err := RegisterSubstance(subs); err != nil {
			return err
		}
	}

	return nil
}
