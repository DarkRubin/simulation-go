package creature

import (
	"simulation/internal/occupier"
	"simulation/internal/pos"
)

type Creature interface {
	occupier.Occupier
	HP() int
	Speed() uint
	Move(pos.Position)
	IsTired() bool
	TargetType() occupier.EntityType
	Hurt(uint)
	Heal(int)
	Recover()
}

var Types = [2]occupier.EntityType{occupier.TypeHerbivore, occupier.TypePredator}
