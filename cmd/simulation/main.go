package main

import (
	"simulation/internal/action"
	"simulation/internal/action/init_"
	"simulation/internal/action/turn"
	"simulation/internal/renderer"
	"simulation/internal/wayfinder"
	"simulation/internal/worldmap"
	"simulation/pkg/setting"
	"simulation/pkg/simulation"
)

func main() {
	sets := setting.Default()
	world := worldmap.New(sets.WorldMapHeight, sets.WorldMapWidth)
	render := renderer.NewCliRenderer(int(world.Size() * 2))

	initActions := []action.Action{
		init_.SpawnOccupiersAction{Settings: *sets},
	}
	turnActions := []action.Action{
		turn.RecoverCreatureAction{},
		turn.CreatureEatAction{},
		turn.CreatureMoveAction{Wayfinder: wayfinder.Bfs{}},
		turn.RemoveDeadCreaturesAction{},
		&turn.RespawnOccupiersAction{Settings: *sets},
	}
	var sim = simulation.New(*sets, world, render, initActions, turnActions)

	sim.Start()
}
