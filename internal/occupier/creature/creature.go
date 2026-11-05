package creature

import (
	"simulation/internal/occupier"
	"simulation/internal/pos"
)

type Creature interface {
	occupier.Occupier
	HP() int
	Speed() int
	Move(way []pos.Position)
	TargetType() occupier.EntityType
}

var Types = [2]occupier.EntityType{occupier.TypeHerbivore, occupier.TypePredator}
