package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/GoGamesLab/Materials/pkg/craft"
)

var Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

func burnFuel(machineID craft.MachineID, fuelID craft.MaterialID, amount float32) craft.Inventory {
	machine := craft.GetMachine(machineID)
	material, _ := craft.GetMaterial(fuelID)

	Logger.Info(fmt.Sprintf("Iniciando combustão: %f unidades de %s na máquina %s", amount, material.Name, machine.Name))

	craft.Store(&machine, fuelID, amount)
	for machine.GetInventory().Materials[fuelID] > 0 {
		machine.Update(1)
		Logger.Info(fmt.Sprintf("🔥 [%s] Temperatura: %.2f°C", machine.Name, machine.Heat))
	}
	return machine.GetInventory()
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
	printResidues(burnFuel(craft.CoalBurnerID, craft.CoalID, 10))

	printResidues(burnFuel(craft.CrudeOilBurnerID, craft.CrudeOilID, 10))
}
