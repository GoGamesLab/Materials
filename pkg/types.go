package materials

import (
	"encoding/json"
	"fmt"
)

type ElementID string
type SubstanceID string
type MaterialID string

type SubstanceState uint8

const (
	Solid SubstanceState = iota + 1
	Liquid
	Gas
	Plasma
)

func (s *SubstanceState) UnmarshalJSON(b []byte) error {
	var stateStr string
	if err := json.Unmarshal(b, &stateStr); err != nil {
		return err
	}
	switch stateStr {
	case "Solid":
		*s = Solid
	case "Liquid":
		*s = Liquid
	case "Gas":
		*s = Gas
	case "Plasma":
		*s = Plasma
	default:
		return fmt.Errorf("🧨 Estado desconhecido: %s", stateStr)
	}
	return nil
}

type Element struct {
	ID           ElementID `json:"id"`
	Name         string    `json:"name"`
	Weight       float64   `json:"weight"`
	BoilingPoint float32   `json:"boilingPoint"`
	Volatility   float32   `json:"volatility"`
}

type ChemicalBond struct {
	Element ElementID `json:"element"`
	Amount  int       `json:"amount"`
}

type Substance struct {
	ID           SubstanceID    `json:"id"`
	Name         string         `json:"name"`
	Composition  []ChemicalBond `json:"composition"`
	MeltingPoint float32        `json:"meltingPoint"`
	BoilingPoint float32        `json:"boilingPoint"`
}

type Composite struct {
	Substance SubstanceID `json:"substance"`
	Quantity  float32     `json:"quantity"`
}

type Material struct {
	ID          MaterialID     `json:"id"`
	Name        string         `json:"name"`
	Composites  []Composite    `json:"composites"`
	Units       float32        `json:"units"`
	State       SubstanceState `json:"state"`
	Temperature float32        `json:"temperature"`
	Hardness    float32        `json:"hardness"`
}
