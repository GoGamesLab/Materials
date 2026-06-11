package materials

import (
	"encoding/json"
	"fmt"
	"os"
)

// const (
// 	VacuumID ElementID = iota
// 	HydrogenID
// 	HeliumID
// 	LithiumID
// 	BerylliumID
// 	BoronID
// 	CarbonID
// 	NitrogenID
// 	OxygenID
// 	FluorineID
// 	NeonID
// 	SodiumID
// 	MagnesiumID
// 	AluminiumID
// 	SiliconID
// 	PhosphorusID
// 	SulfurID
// 	ChlorineID
// 	ArgonID
// 	PotassiumID
// 	CalciumID
// 	TitaniumID
// 	ChromiumID
// 	ManganeseID
// 	IronID
// 	CobaltID
// 	NickelID
// 	CopperID
// 	ZincID
// 	GermaniumID
// 	SilverID
// 	TinID
// 	XenonID
// 	TungstenID
// 	PlatinumID
// 	GoldID
// 	LeadID
// 	ThoriumID
// 	UraniumID
// 	PlutoniumID
// 	CesiumID
// )

var (
	Elements = make(map[ElementID]Element)
)

func init() {
	loadElementsFromJSON("../../Grind/assets/resources/elements.json")
	// RegisterElement(Element{VacuumID, "Vacuum", "V", 0, 0, 1.0})
	// // --- Período 1 ---
	// RegisterElement(Element{HydrogenID, "Hydrogen", "H", 1.008, -252, 0.9})
	// RegisterElement(Element{HeliumID, "Helium", "He", 4.0026, -268, 1.0})
	// // --- Período 2 ---
	// RegisterElement(Element{LithiumID, "Lithium", "Li", 6.94, 1342, 0.2})
	// RegisterElement(Element{BerylliumID, "Beryllium", "Be", 9.0122, 2471, 0.1})
	// RegisterElement(Element{BoronID, "Boron", "B", 10.81, 3927, 0.05})
	// RegisterElement(Element{CarbonID, "Carbon", "C", 12.011, 4827, 0.05})
	// RegisterElement(Element{NitrogenID, "Nitrogen", "N", 14.007, -195, 0.8})
	// RegisterElement(Element{OxygenID, "Oxygen", "O", 15.999, -182, 0.8})
	// RegisterElement(Element{FluorineID, "Fluorine", "F", 18.998, -188, 0.9})
	// RegisterElement(Element{NeonID, "Neon", "Ne", 20.180, -246, 1.0})
	// // --- Período 3 ---
	// RegisterElement(Element{SodiumID, "Sodium", "Na", 22.990, 882, 0.3})
	// RegisterElement(Element{MagnesiumID, "Magnesium", "Mg", 24.305, 1090, 0.2})
	// RegisterElement(Element{AluminiumID, "Aluminium", "Al", 26.982, 2519, 0.1})
	// RegisterElement(Element{SiliconID, "Silicon", "Si", 28.085, 3265, 0.05})
	// RegisterElement(Element{PhosphorusID, "Phosphorus", "P", 30.974, 280, 0.5})
	// RegisterElement(Element{SulfurID, "Sulfur", "S", 32.06, 444, 0.4})
	// RegisterElement(Element{ChlorineID, "Chlorine", "Cl", 35.45, -34, 0.8})
	// RegisterElement(Element{ArgonID, "Argon", "Ar", 39.948, -185, 1.0})
	// // --- Período 4 ---
	// RegisterElement(Element{PotassiumID, "Potassium", "K", 39.098, 759, 0.3})
	// RegisterElement(Element{CalciumID, "Calcium", "Ca", 40.078, 1484, 0.2})
	// RegisterElement(Element{TitaniumID, "Titanium", "Ti", 47.867, 3287, 0.05})
	// RegisterElement(Element{ChromiumID, "Chromium", "Cr", 51.996, 2671, 0.1})
	// RegisterElement(Element{ManganeseID, "Manganese", "Mn", 54.938, 2061, 0.1})
	// RegisterElement(Element{IronID, "Iron", "Fe", 55.845, 2862, 0.05})
	// RegisterElement(Element{CobaltID, "Cobalt", "Co", 58.933, 2927, 0.05})
	// RegisterElement(Element{NickelID, "Nickel", "Ni", 58.693, 2913, 0.05})
	// RegisterElement(Element{CopperID, "Copper", "Cu", 63.546, 2562, 0.05})
	// RegisterElement(Element{ZincID, "Zinc", "Zn", 65.38, 907, 0.4})
	// RegisterElement(Element{GermaniumID, "Germanium", "Ge", 72.63, 2833, 0.1})
	// // --- Outros Importantes ---
	// RegisterElement(Element{SilverID, "Silver", "Ag", 107.87, 2162, 0.1})
	// RegisterElement(Element{TinID, "Tin", "Sn", 118.71, 2602, 0.1})
	// RegisterElement(Element{XenonID, "Xenon", "Xe", 131.29, -108, 1.0})
	// RegisterElement(Element{TungstenID, "Tungsten", "W", 183.84, 5555, 0.01})
	// RegisterElement(Element{PlatinumID, "Platinum", "Pt", 195.08, 3825, 0.05})
	// RegisterElement(Element{GoldID, "Gold", "Au", 196.97, 2856, 0.05})
	// RegisterElement(Element{LeadID, "Lead", "Pb", 207.2, 1749, 0.2})
	// // --- Nucleares ---
	// RegisterElement(Element{ThoriumID, "Thorium", "Th", 232.04, 4788, 0.05})
	// RegisterElement(Element{UraniumID, "Uranium", "U", 238.03, 4131, 0.05})
	// RegisterElement(Element{PlutoniumID, "Plutonium", "Pu", 244, 3228, 0.05})
	// RegisterElement(Element{CesiumID, "Cesium", "Cs", 132.9, 671, 0.6})
}

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

func loadElementsFromJSON(elementsPath string) error {
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
