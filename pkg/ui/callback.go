package ui

import "simulation/internal/worldmap"

type Callback func(worldMap worldmap.Map, turnNumber int, paused bool)
