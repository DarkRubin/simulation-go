package simulation

import (
	"simulation/internal/action"
	"simulation/internal/occupier"
	"simulation/internal/worldmap"
	"simulation/pkg/setting"
	"simulation/pkg/ui"
	"sync/atomic"
	"time"
)

type SingleThreadSimulation struct {
	settings    setting.Settings
	worldMap    *worldmap.Map
	occupiers   occupier.Occupiers
	uiCallback  ui.Callback
	initActions []action.Action
	turnActions []action.Action
	paused      atomic.Bool
	started     atomic.Bool
	turnCounter int
}

func New(s setting.Settings, m *worldmap.Map, callback ui.Callback, init []action.Action, turn []action.Action) Simulation {
	return &SingleThreadSimulation{
		settings:    s,
		worldMap:    m,
		occupiers:   occupier.Occupiers{},
		initActions: init,
		uiCallback:  callback,
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
	s.uiCallback(*s.worldMap, s.turnCounter, s.paused.Load())
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

	s.uiCallback(*s.worldMap, s.turnCounter, s.paused.Load())
	for {
		for s.paused.Load() == true {
			time.Sleep(1000)
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
