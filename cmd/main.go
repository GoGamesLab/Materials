package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/GoGamesLab/Inventory/pkg/container"
	"github.com/GoGamesLab/Materials/pkg/craft"
)

var Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

func burnFuel(machineID craft.MachineID, fuelID craft.MaterialID, amount float32) craft.Supply {
	machine := craft.GetMachine(machineID)
	material, _ := craft.GetMaterial(fuelID)

	Logger.Info(fmt.Sprintf("Iniciando combustão: %f unidades de %s na máquina %s", amount, material.Name, machine.Name))

	craft.Store(&machine, fuelID, amount)
	for machine.GetSupply().Materials.Items[container.ItemID(fuelID)] > 0 {
		machine.Update(1)
		Logger.Info(fmt.Sprintf("🔥 [%s] Temperatura: %.2f°C", machine.Name, machine.Heat))
	}
	return machine.GetSupply()
}

func printResidues(i craft.Supply) {
	for id, j := range i.Materials.Items {
		m, _ := craft.GetMaterial(craft.MaterialID(id))
		Logger.Info(fmt.Sprintf("📦 Produzido %f material %s", j, m.Name))
	}
	for id, j := range i.Substances.Items {
		s, _ := craft.GetSubstance(craft.SubstanceID(id))
		Logger.Info(fmt.Sprintf("🧪 Produzido %f substância %s", j, s.Name))
	}
	for id, j := range i.Elements.Items {
		e, _ := craft.GetElement(craft.ElementID(id))
		Logger.Info(fmt.Sprintf("⚛️️ Produzido %f elemento %s", j, e.Name))
	}
}

func main() {
	printResidues(burnFuel(craft.CoalBurnerID, craft.CoalID, 10))

	printResidues(burnFuel(craft.CrudeOilBurnerID, craft.CrudeOilID, 10))
}
