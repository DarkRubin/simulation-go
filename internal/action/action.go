package action

import (
	"simulation/internal/occupier"
	"simulation/internal/worldmap"
)

type Action interface {
	Act(w *worldmap.Map, m occupier.Occupiers)
}
