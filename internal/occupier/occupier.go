package occupier

import (
	"simulation/internal/pos"
)

type Occupier interface {
	Type() EntityType
	Pos() pos.Position
}

type Occupiers map[EntityType][]Occupier
