package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/GoGamesLab/Materials/pkg/craft"
)

var Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

func burnCoal() craft.Inventory {
	var fogueira craft.Machine = craft.GetMachine(craft.CoalBurnerID)
	Logger.Info("Queimar 10 unidades de Carvão numa fogueira")
	fogueira.AddMaterialToInventory(craft.CoalID, 10)
	for fogueira.GetInventory().Materials[craft.CoalID] > 0 {
		fogueira.Update(1) // Queimar uma unidade por vez!
		Logger.Info(fmt.Sprintf("🔥 Temperatura %f", fogueira.Heat))
	}
	return fogueira.GetInventory()
}

func burnOil() craft.Inventory {
	var tamborOleo craft.Machine = craft.GetMachine(craft.CrudeOilBurnerID)
	Logger.Info("Queimar 10 unidades de Óleo Crú num tambor")
	tamborOleo.AddMaterialToInventory(craft.CrudeOilID, 10)
	for tamborOleo.GetInventory().Materials[craft.CrudeOilID] > 0 {
		tamborOleo.Update(1) // Queimar uma unidade por vez!
		Logger.Info(fmt.Sprintf("🔥 Temperatura %f", tamborOleo.Heat))
	}
	return tamborOleo.GetInventory()
}

func printResidues(i craft.Inventory) {
	for id, j := range i.Materials {
		m, _ := craft.GetMaterial(id)
		if m != nil {
			Logger.Info(fmt.Sprintf("📦 Produzido %f material %s", j, m.Name))
		}
	}
	for id, j := range i.Substances {
		s, _ := craft.GetSubstance(id)
		if s != nil {
			Logger.Info(fmt.Sprintf("🧪 Produzido %f substância %s", j, s.Name))
		}
	}
	for id, j := range i.Elements {
		s, _ := craft.GetElement(id)
		if s != nil {
			Logger.Info(fmt.Sprintf("⚛️️ Produzido %f elemento %s", j, s.Name))
		}
	}
}

func main() {
	printResidues(burnCoal())
	//printResidues(burnOil())
}
