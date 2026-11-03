package simulation

import (
	"simulation/internal/worldmap"
)

type Simulation interface {
	Map() *worldmap.Map
	NextTurn()
	Start()
	Pause()
	Resume()
}
