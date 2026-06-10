package materials

// ID único para busca rápida
type ElementID uint16
type SubstanceID uint16
type MaterialID uint16

type Element struct {
	ID           ElementID
	Name         string
	Symbol       string
	Weight       float64
	BoilingPoint float32 // Temperatura (ºC) em que ele vira gás/sai do material
	Volatility   float32 // Fator de 0.0 a 1.0 (quão rápido ele escapa ao ferver)
}

// Representa a proporção química (ex: H2O -> H:2, O:1)
type ChemicalBond struct {
	Element ElementID
	Amount  int
}

type Substance struct {
	ID           SubstanceID
	Name         string
	composition  []ChemicalBond
	meltingPoint float32
	boilingPoint float32
}

type Composite struct {
	Substance  SubstanceID
	Percentual float32
}

type Material struct {
	ID          MaterialID // 2 bytes - Aponta para MaterialDefinition
	Name        string
	composites  []Composite
	State       SubstanceState // 1 byte  - Estado atual do bloco
	Temperature float32        // 2 bytes - Temperatura de instância para simular mudanças de estado
	Hardness    float32        // Para mecânica de mineração
}
