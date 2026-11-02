package occupier

import (
	"simulation/internal/positon"
)

type Tree struct {
	position positon.Position
}

func (t Tree) Type() EntityType {
	return TypeTree
}

func (t Tree) Pos() positon.Position {
	return t.position
}

func newTree(pos positon.Position) Occupier {
	return Tree{position: pos}
}
