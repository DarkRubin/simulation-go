package creature

import (
	"fmt"
	"simulation/internal/occupier"
	"simulation/internal/pos"
)

func NewCreature(creatType occupier.EntityType, pos pos.Position, hp int, speed int) Creature {
	switch creatType {
	case occupier.TypeHerbivore:
		return NewHerbivore(pos, hp, speed)
	case occupier.TypePredator:
		return NewPredator(pos, 0, hp, speed)
	default:
		panic(fmt.Sprintf("no Creature found for entity type: %s", creatType))
	}
}
