package occupier

import (
	"simulation/internal/positon"
)

type Rock struct {
	position positon.Position
}

func (r Rock) Type() EntityType {
	return TypeRock
}

func (r Rock) Pos() positon.Position {
	return r.position
}

func newRock(pos positon.Position) Occupier {
	return Rock{position: pos}
}
