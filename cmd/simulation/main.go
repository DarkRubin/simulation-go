package main

import (
	"simulation/internal/terminal"
	"simulation/pkg/controller"
	"simulation/pkg/setting"
	"simulation/pkg/setup"
	"simulation/pkg/simulation"
)

func main() {
	sets := setting.Load()

	var sim = simulation.New(
		*sets,
		setup.WorldMap(*sets),
		setup.CallbackFunction(*sets),
		setup.InitActions(*sets),
		setup.TurnActions(*sets),
	)

	controller.ListenAndManage(
		terminal.NewStdInListener(),
		setup.UserInterface(),
		sim,
	)
}
