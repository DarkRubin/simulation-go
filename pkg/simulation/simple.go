package simulation

import (
	"simulation/internal/action"
	"simulation/internal/occupier"
	"simulation/internal/renderer"
	"simulation/internal/worldmap"
	"simulation/pkg/setting"
	"sync/atomic"
	"time"
)

type SingleThreadSimulation struct {
	settings    setting.Settings
	worldMap    *worldmap.Map
	occupiers   occupier.Occupiers
	renderer    renderer.Renderer
	initActions []action.Action
	turnActions []action.Action
	paused      atomic.Bool
	started     atomic.Bool
	turnCounter int
}

func New(s setting.Settings, m *worldmap.Map, r renderer.Renderer, init []action.Action, turn []action.Action) Simulation {
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
	for _, act := range s.turnActions {
		act.Act(s.worldMap, s.occupiers)
	}

	s.turnCounter++
	s.renderer.Render(*s.worldMap, s.turnCounter, s.paused.Load())
}

func (s *SingleThreadSimulation) Start() {
	if s.started.Load() {
		return
	}
	for _, act := range s.initActions {
		act.Act(s.worldMap, s.occupiers)
	}
	s.started.Store(true)
	s.paused.Store(true)

	s.renderer.Render(*s.worldMap, 0, s.paused.Load())
	for {
		for s.paused.Load() {
			time.Sleep(100)
		}
		s.NextTurn()
		time.Sleep(s.settings.GetMoveDelayDuration())
	}
}

func (s *SingleThreadSimulation) Pause() {
	s.paused.Store(true)
}

func (s *SingleThreadSimulation) Resume() {
	s.paused.Store(false)
}

func (s *SingleThreadSimulation) IsPaused() bool {
	return s.paused.Load()
}
