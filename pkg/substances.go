package materials

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
// PureHydrogenID SubstanceID = iota
// CarbonDioxideID
// WaterID
// SodiumChlorideID
// CarbonMonoxideID
// HydrogenOxideID
// MethaneID    // Hidrocarboneto leve (Gás)
// ParaffinID   // Hidrocarboneto pesado (Sólido/Ceroso)
// PureSulfurID // Impureza comum (Enxofre puro)
// PureCarbonID // Carbono elementar (Grafite/Diamante)
// PureIronID   // Ferro elementar (minério de ferro)
// PurePhosphorusID
// PureManganeseID
// PureOxygenID
// PureNitrogenID
// PureSiliconID
// PureArgonID
// UraniumTetrafluorideID
// UraniumHexafluorideID
// UraniumDioxideID
// Cesium137ID
// PlutoniumOxideID
// SugarID
)

var (
	Substances = make(map[SubstanceID]Substance)
)

func init() {
	loadSubstancesFromJSON("../../Grind/assets/materials/substances.json")
	// RegisterSubstance(Substance{
	// 	ID:   PureCarbonID,
	// 	Name: "Carbon",
	// 	composition: []ChemicalBond{
	// 		{Element: CarbonID, Amount: 1},
	// 	},
	// 	meltingPoint: -78.5, // Sublima (vira gás direto)
	// 	boilingPoint: -78.5,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   PureHydrogenID,
	// 	Name: "Hydrogen",
	// 	composition: []ChemicalBond{
	// 		{Element: HydrogenID, Amount: 1},
	// 	},
	// 	meltingPoint: -259.1,
	// 	boilingPoint: -252.9,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   PureOxygenID,
	// 	Name: "Oxygen",
	// 	composition: []ChemicalBond{
	// 		{Element: OxygenID, Amount: 1},
	// 	},
	// 	meltingPoint: -218.8,
	// 	boilingPoint: -183.0,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   CarbonDioxideID,
	// 	Name: "Carbon Dioxide",
	// 	composition: []ChemicalBond{
	// 		{Element: CarbonID, Amount: 1},
	// 		{Element: OxygenID, Amount: 2},
	// 	},
	// 	meltingPoint: -78.5, // Sublima (vira gás direto)
	// 	boilingPoint: -78.5,
	// })
	// RegisterSubstance(Substance{
	// 	// Água Oxigenada (H2O2) simplificada
	// 	ID:   HydrogenOxideID,
	// 	Name: "Hydrogen Peroxide",
	// 	composition: []ChemicalBond{
	// 		{Element: HydrogenID, Amount: 2},
	// 		{Element: OxygenID, Amount: 2},
	// 	},
	// 	meltingPoint: -0.4,
	// 	boilingPoint: 150.2,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   CarbonMonoxideID,
	// 	Name: "Carbon Monoxide",
	// 	composition: []ChemicalBond{
	// 		{Element: CarbonID, Amount: 1},
	// 		{Element: OxygenID, Amount: 1},
	// 	},
	// 	meltingPoint: -205.0,
	// 	boilingPoint: -191.5,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   WaterID,
	// 	Name: "Water",
	// 	composition: []ChemicalBond{
	// 		{Element: HydrogenID, Amount: 2},
	// 		{Element: OxygenID, Amount: 1},
	// 	},
	// 	meltingPoint: 0.0,
	// 	boilingPoint: 100.0,
	// })
	// RegisterSubstance(Substance{
	// 	// Sal de Cozinha (NaCl)
	// 	ID:   SodiumChlorideID,
	// 	Name: "Sodium Chloride",
	// 	composition: []ChemicalBond{
	// 		{Element: SodiumID, Amount: 1},
	// 		{Element: ChlorineID, Amount: 1},
	// 	},
	// 	meltingPoint: 801.0,
	// 	boilingPoint: 1465.0,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   MethaneID,
	// 	Name: "Methane",
	// 	composition: []ChemicalBond{
	// 		{Element: CarbonID, Amount: 1},
	// 		{Element: HydrogenID, Amount: 4},
	// 	},
	// 	meltingPoint: -182.5,
	// 	boilingPoint: -161.5,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   ParaffinID,
	// 	Name: "Paraffin",
	// 	composition: []ChemicalBond{
	// 		{Element: CarbonID, Amount: 20},
	// 		{Element: HydrogenID, Amount: 42},
	// 	},
	// 	meltingPoint: 47.0,
	// 	boilingPoint: 370.0,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   PureSulfurID,
	// 	Name: "Sulfur",
	// 	composition: []ChemicalBond{
	// 		{Element: SulfurID, Amount: 1},
	// 	},
	// 	meltingPoint: 115.2,
	// 	boilingPoint: 444.6,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   PurePhosphorusID,
	// 	Name: "Phosphorus",
	// 	composition: []ChemicalBond{
	// 		{Element: PhosphorusID, Amount: 1},
	// 	},
	// 	meltingPoint: 44.1,
	// 	boilingPoint: 280.0,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   PureManganeseID,
	// 	Name: "Manganese",
	// 	composition: []ChemicalBond{
	// 		{Element: ManganeseID, Amount: 1},
	// 	},
	// 	meltingPoint: 1244.0,
	// 	boilingPoint: 2091.0,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   PureIronID,
	// 	Name: "Iron",
	// 	composition: []ChemicalBond{
	// 		{Element: IronID, Amount: 1},
	// 	},
	// 	meltingPoint: 1538.0,
	// 	boilingPoint: 2862.0,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   PureNitrogenID,
	// 	Name: "Nitrogen",
	// 	composition: []ChemicalBond{
	// 		{Element: NitrogenID, Amount: 1},
	// 	},
	// })
	// RegisterSubstance(Substance{
	// 	ID:   PureSiliconID,
	// 	Name: "Silicon",
	// 	composition: []ChemicalBond{
	// 		{Element: SiliconID, Amount: 1},
	// 	},
	// })
	// RegisterSubstance(Substance{
	// 	ID:   PureArgonID,
	// 	Name: "Argon",
	// 	composition: []ChemicalBond{
	// 		{Element: ArgonID, Amount: 1},
	// 	},
	// })
	// RegisterSubstance(Substance{
	// 	ID:   SugarID,
	// 	Name: "Sugar",
	// 	composition: []ChemicalBond{
	// 		{CarbonID, 6},
	// 		{HydrogenID, 12},
	// 		{OxygenID, 6},
	// 	},
	// })
	// RegisterSubstance(Substance{
	// 	ID:   UraniumTetrafluorideID,
	// 	Name: "Uranium Tetrafluoride",
	// 	composition: []ChemicalBond{
	// 		{Element: UraniumID, Amount: 1},
	// 		{Element: FluorineID, Amount: 4},
	// 	},
	// 	meltingPoint: 960.0,
	// 	boilingPoint: 1417.0,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   UraniumHexafluorideID,
	// 	Name: "Uranium Hexafluoride",
	// 	composition: []ChemicalBond{
	// 		{Element: UraniumID, Amount: 1},
	// 		{Element: FluorineID, Amount: 6},
	// 	},
	// 	meltingPoint: 64.0,
	// 	boilingPoint: 56.5, // Sublima a pressões normais
	// })
	// RegisterSubstance(Substance{
	// 	ID:   UraniumDioxideID,
	// 	Name: "Uranium Dioxide",
	// 	composition: []ChemicalBond{
	// 		{Element: UraniumID, Amount: 1},
	// 		{Element: OxygenID, Amount: 2},
	// 	},
	// 	meltingPoint: 2865.0,
	// 	boilingPoint: 3500.0,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   Cesium137ID,
	// 	Name: "Cesium-137 Slurry",
	// 	composition: []ChemicalBond{
	// 		{Element: CesiumID, Amount: 1},
	// 	},
	// 	meltingPoint: 28.5,
	// 	boilingPoint: 671.0,
	// })
	// RegisterSubstance(Substance{
	// 	ID:   PlutoniumOxideID,
	// 	Name: "Plutonium Dioxide",
	// 	composition: []ChemicalBond{
	// 		{Element: PlutoniumID, Amount: 1},
	// 		{Element: OxygenID, Amount: 2},
	// 	},
	// 	meltingPoint: 2400.0,
	// 	boilingPoint: 2800.0,
	// })
}

func RegisterSubstance(s Substance) error {
	if _, exists := Substances[s.ID]; exists {
		return fmt.Errorf("🧨 Substância com ID %d já registrada", s.ID)
	}
	Substances[s.ID] = s

	return nil
}

func GetSubstance(id SubstanceID) (*Substance, error) {
	if s, ok := Substances[id]; ok {
		return &s, nil
	}
	return nil, fmt.Errorf("🧨 Substance %v: not found", id)
}

func (s Substance) GetMolecularWeight() float64 {
	var total float64
	for _, bond := range s.Composition {
		// Busca o elemento no registro pelo ID (Número Atômico)
		element := Elements[bond.Element]
		total += element.Weight * float64(bond.Amount)
	}
	return total
}

type State int

func (s Substance) GetState(currentTemp float32) SubstanceState {
	// Precisamos definir também o MeltingPoint (Ponto de Fusão) na Substance
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

func loadSubstancesFromJSON(substancesPath string) error {
	// Carregar Substâncias
	sData, err := os.ReadFile(substancesPath)
	if err != nil {
		return fmt.Errorf("erro lendo substâncias: %w", err)
	}
	var sList []Substance
	if err := json.Unmarshal(sData, &sList); err != nil {
		return err
	}
	for _, s := range sList {
		Substances[s.ID] = s
	}

	return nil
}
