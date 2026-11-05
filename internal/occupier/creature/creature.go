package creature

import (
	"simulation/internal/occupier"
	"simulation/internal/pos"
)

type Creature interface {
	occupier.Occupier
	HP() int
	Speed() uint
	Move(way []pos.Position) bool
	Moved() bool
	TargetType() occupier.EntityType
	Eat(occupier.Occupier) bool
	Recover()
}

var Types = [2]occupier.EntityType{occupier.TypeHerbivore, occupier.TypePredator}
