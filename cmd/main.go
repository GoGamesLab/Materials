package main

import (
	"fmt"

	"github.com/GoGamesLab/Materials/pkg/craft"
)

func main() {
	var fogueira craft.Machine = craft.GetMachine(craft.CoalBurnerID)
	fmt.Println("Queimar 10 unidades de Carvão numa fogueira")
	fogueira.AddMaterialToInventory(craft.CoalID, 10)
	for fogueira.GetInventory().Materials[craft.CoalID] > 0 {
		fogueira.Update(1) // Queimar uma unidade por vez!
		fmt.Printf("Temperatura %f\n", fogueira.Heat)
	}
	fogueiraInv := fogueira.GetInventory()
	for id, j := range fogueiraInv.Materials {
		m, _ := craft.GetMaterial(id)
		if m != nil {
			fmt.Printf("📦 Produzido %f material %s\n", j, m.Name)
		}
	}
	for id, j := range fogueiraInv.Substances {
		s, _ := craft.GetSubstance(id)
		if s != nil {
			fmt.Printf("🧪 Produzido %f substância %s\n", j, s.Name)
		}
	}
	for id, j := range fogueiraInv.Elements {
		s, _ := craft.GetElement(id)
		if s != nil {
			fmt.Printf("🧪 Produzido %f elemento %s\n", j, s.Name)
		}
	}

	// var tamborOleo craft.Machine = craft.GetMachine(craft.CrudeOilBurnerID)
	// fmt.Println("Queimar 10 unidades de Óleo Crú num tambor")
	// tamborOleo.AddMaterialToInventory(craft.CrudeOilID, 10)
	// for tamborOleo.InternalFuel > 0 {
	// 	tamborOleo.Update(1) // Queimar uma unidade por vez!
	// 	fmt.Printf("Combustível restante %f. Temperatura %f\n", tamborOleo.InternalFuel, tamborOleo.Heat)
	// }
	// tamborInv := tamborOleo.GetInventory()
	// for id, j := range tamborInv.Materials {
	// 	m := craft.GetMaterial(id)
	// 	if m != nil {
	// 		fmt.Printf("📦 Produzido %f material %s\n", j, m.Name)
	// 	}
	// }
	// for id, j := range tamborInv.Substances {
	// 	s := craft.GetSubstance(id)
	// 	if s != nil {
	// 		fmt.Printf("🧪 Produzido %f substância %s\n", j, s.Name)
	// 	}
	// }

}
