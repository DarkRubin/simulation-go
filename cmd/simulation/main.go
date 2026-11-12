package main

import (
	"simulation/internal/action"
	"simulation/internal/action/init_"
	"simulation/internal/action/turn"
	"simulation/internal/occupier/spawner"
	"simulation/internal/renderer"
	"simulation/internal/wayfinder"
	"simulation/internal/worldmap"
	"simulation/pkg/command"
	"simulation/pkg/controller"
	"simulation/pkg/setting"
	"simulation/pkg/simulation"
)

func main() {
	sets := setting.Default()
	world := worldmap.New(sets.WorldMapHeight, sets.WorldMapWidth)
	summoner := spawner.Spawner{Settings: *sets}
	render := renderer.NewTuiRenderer(int(world.Size() * 2))

	initActions := []action.Action{
		init_.SpawnOccupiersAction{Spawner: summoner},
	}
	turnActions := []action.Action{
		turn.RecoverCreatureAction{},
		turn.HerbivoresEatAction{},
		turn.PredatorsAttackAction{},
		turn.CreatureMoveAction{Wayfinder: wayfinder.Bfs{}},
		turn.HungerCreatureAction{},
		turn.RemoveDeadCreaturesAction{},
		turn.RespawnOccupiersAction{Spawner: summoner},
	}
	var sim = simulation.New(*sets, world, render, initActions, turnActions)
	go sim.Start()

	controller.ListenAndManage(command.NewKeyboardCommandReader(), sim)
}
