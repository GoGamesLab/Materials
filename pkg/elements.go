package materials

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Elements = make(map[ElementID]Element)
)

func RegisterElement(e Element) error {
	if _, exists := Elements[e.ID]; exists {
		return fmt.Errorf("🧨 Elemento com ID %s já registrado", e.ID)
	}

	Elements[e.ID] = e

	return nil
}

func GetElement(id ElementID) (*Element, error) {
	if e, ok := Elements[id]; ok {
		return &e, nil
	}
	return nil, fmt.Errorf("🧨 Element %v: not found", id)
}

func LoadElementsFromJSON(elementsPath string) error {
	// Carregar Elementos
	eData, err := os.ReadFile(elementsPath)
	if err != nil {
		return fmt.Errorf("erro lendo elementos: %w", err)
	}
	var eList []Element
	if err := json.Unmarshal(eData, &eList); err != nil {
		return err
	}
	for _, e := range eList {
		Elements[e.ID] = e
	}

	return nil
}
