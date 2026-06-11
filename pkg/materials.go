package materials

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	Materials  = make(map[MaterialID]Material)
	Signatures = make(map[string][]MaterialID)
)

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

func LoadMaterialsFromJSON(materialsPath string) error {

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
