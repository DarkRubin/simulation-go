package occupier

import (
	"simulation/internal/pos"
)

type Grass struct {
	position pos.Position
}

func (g *Grass) Type() EntityType {
	return TypeGrass
}

func (g *Grass) Pos() pos.Position {
	return g.position
}

func newGrass(pos pos.Position) Occupier {
	return &Grass{position: pos}
}
