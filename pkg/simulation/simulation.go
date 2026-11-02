package simulation

import (
	"simulation/internal/worldmap"
	"time"
)

type Simulation interface {
	Map() *worldmap.Map
	NextTurn()
	Start(time.Duration)
	Pause()
	Resume()
}
