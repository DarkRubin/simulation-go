package occupier

import (
	"simulation/internal/pos"
)

type Rock struct {
	position pos.Position
}

func (r *Rock) Type() EntityType {
	return TypeRock
}

func (r *Rock) Pos() pos.Position {
	return r.position
}

func NewRock(pos pos.Position) Occupier {
	return &Rock{position: pos}
}
