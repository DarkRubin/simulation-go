package simulation

import (
	"fmt"
	"simulation/internal/action"
	"simulation/internal/occupier"
	"simulation/internal/renderer"
	"simulation/internal/settings"
	"simulation/internal/worldmap"
	"sync/atomic"
	"time"
)

type SingleThreadSimulation struct {
	settings    settings.Settings
	worldMap    *worldmap.Map
	occupiers   occupier.Occupiers
	renderer    renderer.Renderer
	initActions []action.Action
	turnActions []action.Action
	paused      atomic.Bool
	started     atomic.Bool
	turnCounter int
}

func New(s settings.Settings, m *worldmap.Map, r renderer.Renderer, init []action.Action, turn []action.Action) Simulation {
	return &SingleThreadSimulation{
		settings:    s,
		worldMap:    m,
		occupiers:   occupier.Occupiers{},
		renderer:    r,
		initActions: init,
		turnActions: turn,
	}
}

func (s *SingleThreadSimulation) Map() *worldmap.Map {
	return s.worldMap
}

func (s *SingleThreadSimulation) NextTurn() {
	if !s.started.Load() {
		panic("Simulation haven't started")
	}

	for _, act := range s.turnActions {
		act.Act(s.worldMap, s.occupiers)
	}

	s.turnCounter++
	fmt.Println("Move ", s.turnCounter)
	s.renderer.Render(*s.worldMap)
}

func (s *SingleThreadSimulation) Start() {
	for _, act := range s.initActions {
		act.Act(s.worldMap, s.occupiers)
	}
	s.started.Store(true)

	fmt.Println("Original simulation")
	s.renderer.Render(*s.worldMap)

	for !s.paused.Load() {
		s.NextTurn()
		time.Sleep(s.settings.MoveDelay())
	}
}

func (s *SingleThreadSimulation) Pause() {
	s.paused.Store(true)
}

func (s *SingleThreadSimulation) Resume() {
	s.paused.Store(false)
}
