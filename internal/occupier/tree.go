package occupier

import (
	"simulation/internal/pos"
)

type Tree struct {
	position pos.Position
}

func (t *Tree) Type() EntityType {
	return TypeTree
}

func (t *Tree) Pos() pos.Position {
	return t.position
}

func NewTree(pos pos.Position) Occupier {
	return &Tree{position: pos}
}
