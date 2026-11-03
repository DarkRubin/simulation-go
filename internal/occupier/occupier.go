package occupier

import (
	"simulation/internal/pos"
)

type Occupier interface {
	Type() EntityType
	Pos() pos.Position
}

func (e EntityType) String() string {
	return EntityTypeString[e]
}
