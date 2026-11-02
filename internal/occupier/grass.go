package occupier

import (
	"simulation/internal/positon"
)

type Grass struct {
	position positon.Position
}

func (g Grass) Type() EntityType {
	return TypeGrass
}

func (g Grass) Pos() positon.Position {
	return g.position
}

func newGrass(pos positon.Position) Occupier {
	return Grass{position: pos}
}
