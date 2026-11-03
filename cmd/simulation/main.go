package main

import (
	"simulation/internal/action"
	"simulation/internal/renderer"
	"simulation/internal/wayfinder"
	"simulation/internal/worldmap"
	"simulation/pkg/simulation"
	"time"
)

const (
	DefaultMoveDelay = 400 * time.Millisecond
	DefaultMapHeight = 6
	DefaultMapWidth  = 7
)

func main() {
	world := worldmap.New(DefaultMapHeight, DefaultMapWidth)
	render := renderer.NewCliRenderer(world)

	initActions := []action.Action{action.PutEntitiesInitAction{}}
	turnActions := []action.Action{action.MoveOrEatAction{Wayfinder: wayfinder.Bfs{}}}
	var sim = simulation.New(world, render, initActions, turnActions)

	sim.Start(DefaultMoveDelay)
}
