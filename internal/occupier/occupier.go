package occupier

import (
	"simulation/internal/positon"
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
	Pos() positon.Position
}

func (e EntityType) IsCreature() bool {
	return e == TypePredator || e == TypeHerbivore
}

func (e EntityType) String() string {
	return EntityTypeString[e]
}
