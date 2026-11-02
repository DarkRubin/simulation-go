package simulation

import (
	"fmt"
	"simulation/internal/action"
	"simulation/internal/renderer"
	"simulation/internal/worldmap"
	"sync/atomic"
	"time"
)

type SingleThreadSimulation struct {
	worldMap    *worldmap.Map
	renderer    renderer.Renderer
	initActions []action.Action
	turnActions []action.Action
	paused      atomic.Bool
	started     atomic.Bool
	turnCounter int
}

func New(m *worldmap.Map, r renderer.Renderer, init []action.Action, turn []action.Action) Simulation {
	return &SingleThreadSimulation{
		worldMap:    m,
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
		act.Act(s.worldMap)
	}
	s.turnCounter++
	fmt.Println("Move ", s.turnCounter)
	s.renderer.Render(*s.worldMap)
}

func (s *SingleThreadSimulation) Start(moveDelay time.Duration) {
	for _, act := range s.initActions {
		act.Act(s.worldMap)
	}
	s.started.Store(true)

	for !s.paused.Load() {
		s.NextTurn()
		time.Sleep(moveDelay)
	}
}

func (s *SingleThreadSimulation) Pause() {
	s.paused.Store(true)
}

func (s *SingleThreadSimulation) Resume() {
	s.paused.Store(false)
}
