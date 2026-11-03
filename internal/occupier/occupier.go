package occupier

import (
	"simulation/internal/pos"
)

const (
	TypeGrass EntityType = iota
	TypeRock
	TypeTree
	TypeHerbivore
	TypePredator
	TypeCount_
)

var EntityTypeString = map[EntityType]string{
	TypeGrass:     "Grass",
	TypeRock:      "Rock",
	TypeTree:      "Tree",
	TypeHerbivore: "Herbivore",
	TypePredator:  "Predator",
}

type EntityType int

type Occupier interface {
	Type() EntityType
	Pos() pos.Position
}

func (e EntityType) String() string {
	return EntityTypeString[e]
}
