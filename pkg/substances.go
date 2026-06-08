package materials

import (
	"fmt"
)

type SubstanceState uint8

const (
	PureHydrogenID SubstanceID = iota
	CarbonDioxideID
	WaterID
	SodiumChlorideID
	CarbonMonoxideID
	HydrogenOxideID
	MethaneID    // Hidrocarboneto leve (Gás)
	ParaffinID   // Hidrocarboneto pesado (Sólido/Ceroso)
	PureSulfurID // Impureza comum (Enxofre puro)
	PureCarbonID // Carbono elementar (Grafite/Diamante)
	PureIronID   // Ferro elementar (minério de ferro)
	PureOxygenID
	PureNitrogenID
	PureSiliconID
	PureArgonID
	UraniumTetrafluorideID
	UraniumHexafluorideID
	UraniumDioxideID
	Cesium137ID
	PlutoniumOxideID
	SugarID
)

var (
	Substances = make(map[SubstanceID]Substance)
)

func init() {
	RegisterSubstance(Substance{
		ID:   PureCarbonID,
		Name: "Carbon",
		Composition: []ChemicalBond{
			{Element: CarbonID, Amount: 1},
		},
		MeltingPoint: -78.5, // Sublima (vira gás direto)
		BoilingPoint: -78.5,
	})
	RegisterSubstance(Substance{
		ID:   PureHydrogenID,
		Name: "Hydrogen",
		Composition: []ChemicalBond{
			{Element: HydrogenID, Amount: 1},
		},
		MeltingPoint: -259.1,
		BoilingPoint: -252.9,
	})
	RegisterSubstance(Substance{
		ID:   PureOxygenID,
		Name: "Oxygen",
		Composition: []ChemicalBond{
			{Element: OxygenID, Amount: 1},
		},
		MeltingPoint: -218.8,
		BoilingPoint: -183.0,
	})
	RegisterSubstance(Substance{
		ID:   CarbonDioxideID,
		Name: "Carbon Dioxide",
		Composition: []ChemicalBond{
			{Element: CarbonID, Amount: 1},
			{Element: OxygenID, Amount: 2},
		},
		MeltingPoint: -78.5, // Sublima (vira gás direto)
		BoilingPoint: -78.5,
	})
	RegisterSubstance(Substance{
		// Água Oxigenada (H2O2) simplificada
		ID:   HydrogenOxideID,
		Name: "Hydrogen Peroxide",
		Composition: []ChemicalBond{
			{Element: HydrogenID, Amount: 2},
			{Element: OxygenID, Amount: 2},
		},
		MeltingPoint: -0.4,
		BoilingPoint: 150.2,
	})
	RegisterSubstance(Substance{
		ID:   CarbonMonoxideID,
		Name: "Carbon Monoxide",
		Composition: []ChemicalBond{
			{Element: CarbonID, Amount: 1},
			{Element: OxygenID, Amount: 1},
		},
		MeltingPoint: -205.0,
		BoilingPoint: -191.5,
	})
	RegisterSubstance(Substance{
		ID:   WaterID,
		Name: "Water",
		Composition: []ChemicalBond{
			{Element: HydrogenID, Amount: 2},
			{Element: OxygenID, Amount: 1},
		},
		MeltingPoint: 0.0,
		BoilingPoint: 100.0,
	})
	RegisterSubstance(Substance{
		// Sal de Cozinha (NaCl)
		ID:   SodiumChlorideID,
		Name: "Sodium Chloride",
		Composition: []ChemicalBond{
			{Element: SodiumID, Amount: 1},
			{Element: ChlorineID, Amount: 1},
		},
		MeltingPoint: 801.0,
		BoilingPoint: 1465.0,
	})
	RegisterSubstance(Substance{
		ID:   MethaneID,
		Name: "Methane",
		Composition: []ChemicalBond{
			{Element: CarbonID, Amount: 1},
			{Element: HydrogenID, Amount: 4},
		},
		MeltingPoint: -182.5,
		BoilingPoint: -161.5,
	})
	RegisterSubstance(Substance{
		ID:   ParaffinID,
		Name: "Paraffin",
		Composition: []ChemicalBond{
			{Element: CarbonID, Amount: 20},
			{Element: HydrogenID, Amount: 42},
		},
		MeltingPoint: 47.0,
		BoilingPoint: 370.0,
	})
	RegisterSubstance(Substance{
		ID:   PureSulfurID,
		Name: "Sulfur",
		Composition: []ChemicalBond{
			{Element: SulfurID, Amount: 1},
		},
		MeltingPoint: 115.2,
		BoilingPoint: 444.6,
	})
	RegisterSubstance(Substance{
		ID:   PureIronID,
		Name: "Iron",
		Composition: []ChemicalBond{
			{Element: IronID, Amount: 1},
		},
		MeltingPoint: 1538.0,
		BoilingPoint: 2862.0,
	})
	RegisterSubstance(Substance{
		ID:   PureNitrogenID,
		Name: "Nitrogen",
		Composition: []ChemicalBond{
			{Element: NitrogenID, Amount: 1},
		},
	})
	RegisterSubstance(Substance{
		ID:   PureSiliconID,
		Name: "Silicon",
		Composition: []ChemicalBond{
			{Element: SiliconID, Amount: 1},
		},
	})
	RegisterSubstance(Substance{
		ID:   PureArgonID,
		Name: "Argon",
		Composition: []ChemicalBond{
			{Element: ArgonID, Amount: 1},
		},
	})
	RegisterSubstance(Substance{
		ID:   SugarID,
		Name: "Sugar",
		Composition: []ChemicalBond{
			{CarbonID, 6},
			{HydrogenID, 12},
			{OxygenID, 6},
		},
	})
	RegisterSubstance(Substance{
		ID:   UraniumTetrafluorideID,
		Name: "Uranium Tetrafluoride",
		Composition: []ChemicalBond{
			{Element: UraniumID, Amount: 1},
			{Element: FluorineID, Amount: 4},
		},
		MeltingPoint: 960.0,
		BoilingPoint: 1417.0,
	})
	RegisterSubstance(Substance{
		ID:   UraniumHexafluorideID,
		Name: "Uranium Hexafluoride",
		Composition: []ChemicalBond{
			{Element: UraniumID, Amount: 1},
			{Element: FluorineID, Amount: 6},
		},
		MeltingPoint: 64.0,
		BoilingPoint: 56.5, // Sublima a pressões normais
	})
	RegisterSubstance(Substance{
		ID:   UraniumDioxideID,
		Name: "Uranium Dioxide",
		Composition: []ChemicalBond{
			{Element: UraniumID, Amount: 1},
			{Element: OxygenID, Amount: 2},
		},
		MeltingPoint: 2865.0,
		BoilingPoint: 3500.0,
	})
	RegisterSubstance(Substance{
		ID:   Cesium137ID,
		Name: "Cesium-137 Slurry",
		Composition: []ChemicalBond{
			{Element: CesiumID, Amount: 1},
		},
		MeltingPoint: 28.5,
		BoilingPoint: 671.0,
	})
	RegisterSubstance(Substance{
		ID:   PlutoniumOxideID,
		Name: "Plutonium Dioxide",
		Composition: []ChemicalBond{
			{Element: PlutoniumID, Amount: 1},
			{Element: OxygenID, Amount: 2},
		},
		MeltingPoint: 2400.0,
		BoilingPoint: 2800.0,
	})
}

func RegisterSubstance(s Substance) error {
	if _, exists := Substances[s.ID]; exists {
		return fmt.Errorf("🧨 Substância com ID %d já registrada", s.ID)
	}
	Substances[s.ID] = s

	return nil
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

func GetSubstance(id SubstanceID) (*Substance, error) {
	if s, ok := Substances[id]; ok {
		return &s, nil
	}
	return nil, fmt.Errorf("🧨 Substance %v: not found", id)
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
