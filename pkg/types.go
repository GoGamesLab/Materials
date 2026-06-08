package materials

// ID único para busca rápida
type ElementID uint16
type SubstanceID uint16
type MaterialID uint16

type Element struct {
	ID            ElementID
	Name          string
	Symbol        string
	Weight        float64
	Category      string
	IsRadioactive bool    // Para mecânicas de dano por radiação e fissão
	FusionFuel    bool    // Se pode ser usado em reatores de fusão (H, He, Li)
	BoilingPoint  float32 // Temperatura (ºC) em que ele vira gás/sai do material
	Volatility    float32 // Fator de 0.0 a 1.0 (quão rápido ele escapa ao ferver)
}

// Representa a proporção química (ex: H2O -> H:2, O:1)
type ChemicalBond struct {
	Element ElementID
	Amount  int
}

type Substance struct {
	ID           SubstanceID
	Name         string
	Composition  []ChemicalBond
	MeltingPoint float32
	BoilingPoint float32
}

type Composite struct {
	Substance  SubstanceID
	Percentual float32
}

type Material struct {
	ID            MaterialID // 2 bytes - Aponta para MaterialDefinition
	Name          string
	Composites    []Composite
	State         SubstanceState // 1 byte  - Estado atual do bloco
	HP            uint8          // 1 byte  - Integridade (0-255)
	Temperature   float32        // 2 bytes - Temperatura local para simular mudanças de estado
	Hardness      float32        // Para mecânica de mineração
	Density       float32
	SpecificHeat  float32 // Calor específico: quanto de energia precisa para mudar 1°C
	EnergyDensity float32 // "Calorias": Joules por unidade de massa se queimado/reagido
	FlashPoint    float32 // Temperatura em que o material entra em combustão
	Conductivity  float32 // Quão rápido ele transfere calor para vizinhos
}
