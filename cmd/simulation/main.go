package main

import (
	"simulation/internal/action"
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

	initActions := []action.Action{action.PutEntitiesInitAction{}}
	turnActions := []action.Action{action.MoveOrEatAction{Wayfinder: wayfinder.Bfs{}}}
	var sim = simulation.New(*sets, world, render, initActions, turnActions)

	sim.Start()
}
