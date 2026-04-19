package craft

import (
	"fmt"
	"sync"
)

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

type CategoryType int

const (
	Nonmetal CategoryType = iota
	AlkaliMetal
	AlkalineEarthMetal
	PostTransitionMetal
	TransitionMetal
	Metalloid
	Halogen
	NobleGas
	Actinide
)

const (
	VacuumID     ElementID = 0
	HydrogenID   ElementID = 1
	HeliumID     ElementID = 2
	LithiumID    ElementID = 3
	BerylliumID  ElementID = 4
	BoronID      ElementID = 5
	CarbonID     ElementID = 6
	NitrogenID   ElementID = 7
	OxygenID     ElementID = 8
	FluorineID   ElementID = 9
	NeonID       ElementID = 10
	SodiumID     ElementID = 11
	MagnesiumID  ElementID = 12
	AluminiumID  ElementID = 13
	SiliconID    ElementID = 14
	PhosphorusID ElementID = 15
	SulfurID     ElementID = 16
	ChlorineID   ElementID = 17
	ArgonID      ElementID = 18
	PotassiumID  ElementID = 19
	CalciumID    ElementID = 20
	TitaniumID   ElementID = 22
	ChromiumID   ElementID = 24
	ManganeseID  ElementID = 25
	IronID       ElementID = 26
	CobaltID     ElementID = 27
	NickelID     ElementID = 28
	CopperID     ElementID = 29
	ZincID       ElementID = 30
	GermaniumID  ElementID = 32
	SilverID     ElementID = 47
	TinID        ElementID = 50
	XenonID      ElementID = 54
	TungstenID   ElementID = 74
	PlatinumID   ElementID = 78
	GoldID       ElementID = 79
	LeadID       ElementID = 82
	ThoriumID    ElementID = 90
	UraniumID    ElementID = 92
	PlutoniumID  ElementID = 94
)

var (
	Elements      = make(map[ElementID]Element)
	elementsMutex sync.Mutex
)

func init() {
	RegisterElement(Element{VacuumID, "Vacuum", "V", 0, "None", false, false, 0, 1.0})
	// --- Período 1 ---
	RegisterElement(Element{HydrogenID, "Hydrogen", "H", 1.008, "Nonmetal", false, true, -252, 0.9})
	RegisterElement(Element{HeliumID, "Helium", "He", 4.0026, "Noble Gas", false, true, -268, 1.0})
	// --- Período 2 ---
	RegisterElement(Element{LithiumID, "Lithium", "Li", 6.94, "Alkali Metal", false, true, 1342, 0.2})
	RegisterElement(Element{BerylliumID, "Beryllium", "Be", 9.0122, "Alkaline Earth Metal", false, false, 2471, 0.1})
	RegisterElement(Element{BoronID, "Boron", "B", 10.81, "Metalloid", false, false, 3927, 0.05})
	RegisterElement(Element{CarbonID, "Carbon", "C", 12.011, "Nonmetal", false, false, 4827, 0.05})
	RegisterElement(Element{NitrogenID, "Nitrogen", "N", 14.007, "Nonmetal", false, false, -195, 0.8})
	RegisterElement(Element{OxygenID, "Oxygen", "O", 15.999, "Nonmetal", false, false, -182, 0.8})
	RegisterElement(Element{FluorineID, "Fluorine", "F", 18.998, "Halogen", false, false, -188, 0.9})
	RegisterElement(Element{NeonID, "Neon", "Ne", 20.180, "Noble Gas", false, false, -246, 1.0})
	// --- Período 3 ---
	RegisterElement(Element{SodiumID, "Sodium", "Na", 22.990, "Alkali Metal", false, false, 882, 0.3})
	RegisterElement(Element{MagnesiumID, "Magnesium", "Mg", 24.305, "Alkaline Earth Metal", false, false, 1090, 0.2})
	RegisterElement(Element{AluminiumID, "Aluminium", "Al", 26.982, "Post-transition Metal", false, false, 2519, 0.1})
	RegisterElement(Element{SiliconID, "Silicon", "Si", 28.085, "Metalloid", false, false, 3265, 0.05})
	RegisterElement(Element{PhosphorusID, "Phosphorus", "P", 30.974, "Nonmetal", false, false, 280, 0.5})
	RegisterElement(Element{SulfurID, "Sulfur", "S", 32.06, "Nonmetal", false, false, 444, 0.4})
	RegisterElement(Element{ChlorineID, "Chlorine", "Cl", 35.45, "Halogen", false, false, -34, 0.8})
	RegisterElement(Element{ArgonID, "Argon", "Ar", 39.948, "Noble Gas", false, false, -185, 1.0})
	// --- Período 4 ---
	RegisterElement(Element{PotassiumID, "Potassium", "K", 39.098, "Alkali Metal", false, false, 759, 0.3})
	RegisterElement(Element{CalciumID, "Calcium", "Ca", 40.078, "Alkaline Earth Metal", false, false, 1484, 0.2})
	RegisterElement(Element{TitaniumID, "Titanium", "Ti", 47.867, "Transition Metal", false, false, 3287, 0.05})
	RegisterElement(Element{ChromiumID, "Chromium", "Cr", 51.996, "Transition Metal", false, false, 2671, 0.1})
	RegisterElement(Element{ManganeseID, "Manganese", "Mn", 54.938, "Transition Metal", false, false, 2061, 0.1})
	RegisterElement(Element{IronID, "Iron", "Fe", 55.845, "Transition Metal", false, false, 2862, 0.05})
	RegisterElement(Element{CobaltID, "Cobalt", "Co", 58.933, "Transition Metal", false, false, 2927, 0.05})
	RegisterElement(Element{NickelID, "Nickel", "Ni", 58.693, "Transition Metal", false, false, 2913, 0.05})
	RegisterElement(Element{CopperID, "Copper", "Cu", 63.546, "Transition Metal", false, false, 2562, 0.05})
	RegisterElement(Element{ZincID, "Zinc", "Zn", 65.38, "Transition Metal", false, false, 907, 0.4})
	RegisterElement(Element{GermaniumID, "Germanium", "Ge", 72.63, "Metalloid", false, false, 2833, 0.1})
	// --- Outros Importantes ---
	RegisterElement(Element{SilverID, "Silver", "Ag", 107.87, "Transition Metal", false, false, 2162, 0.1})
	RegisterElement(Element{TinID, "Tin", "Sn", 118.71, "Post-transition Metal", false, false, 2602, 0.1})
	RegisterElement(Element{XenonID, "Xenon", "Xe", 131.29, "Noble Gas", false, false, -108, 1.0})
	RegisterElement(Element{TungstenID, "Tungsten", "W", 183.84, "Transition Metal", false, false, 5555, 0.01})
	RegisterElement(Element{PlatinumID, "Platinum", "Pt", 195.08, "Transition Metal", false, false, 3825, 0.05})
	RegisterElement(Element{GoldID, "Gold", "Au", 196.97, "Transition Metal", false, false, 2856, 0.05})
	RegisterElement(Element{LeadID, "Lead", "Pb", 207.2, "Post-transition Metal", false, false, 1749, 0.2})
	// --- Nucleares ---
	RegisterElement(Element{ThoriumID, "Thorium", "Th", 232.04, "Actinide", true, false, 4788, 0.05})
	RegisterElement(Element{UraniumID, "Uranium", "U", 238.03, "Actinide", true, false, 4131, 0.05})
	RegisterElement(Element{PlutoniumID, "Plutonium", "Pu", 244, "Actinide", true, false, 3228, 0.05})
}

func RegisterElement(e Element) error {
	elementsMutex.Lock()
	defer elementsMutex.Unlock()

	if _, exists := Elements[e.ID]; exists {
		return fmt.Errorf("🧨 Elemento com ID %d já registrada", e.ID)
	}

	Elements[e.ID] = e

	fmt.Printf("⚛ Elemento %s registrado\n", e.Name)

	return nil
}

func GetElement(id ElementID) (*Element, error) {
	if e, ok := Elements[id]; ok {
		return &e, nil
	}
	return nil, fmt.Errorf("🧨 Element %v: not found", id)
}
