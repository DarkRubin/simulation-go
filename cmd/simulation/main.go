package main

import (
	"fmt"
	"simulation/internal/action"
	"simulation/internal/action/init_"
	"simulation/internal/action/turn"
	"simulation/internal/renderer"
	"simulation/internal/settings"
	"simulation/internal/wayfinder"
	"simulation/internal/worldmap"
	"simulation/pkg/simulation"
)

func main() {
	sets := settings.Default()
	world := worldmap.New(sets.WorldHeight, sets.WorldWidth)
	render := renderer.NewCliRenderer(int(world.Size() * 2))

	initActions := []action.Action{init_.SpawnOccupiersAction{Settings: sets}}
	turnActions := []action.Action{
		turn.RecoverCreatureAction{},
		turn.CreatureEatAction{},
		turn.CreatureMoveAction{Wayfinder: wayfinder.Bfs{}},
		turn.RemoveDeadCreaturesAction{},
	}
	var sim = simulation.New(*sets, world, render, initActions, turnActions)

	sim.Start()
}
