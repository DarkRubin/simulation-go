package turn

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/worldmap"
)

type CreatureEatAction struct{}

func (c CreatureEatAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	forEachCreature(w, o, func(creature creature.Creature) {
		if creature.Moved() {
			return
		}
		oldPos := creature.Pos()

		neighbors := w.GetAll(creature.Pos().Neighbors())
		food := findNearby(neighbors, creature.TargetType())
		if food != nil && creature.Eat(food) {
			w.Remove(oldPos)
			w.Put(creature.Pos(), creature)
			o.RemovePos(creature.TargetType(), food.Pos())
			o.UpdatePos(creature.Type(), oldPos, creature.Pos())
		}
	})
}

func findNearby(neighbors []occupier.Occupier, target occupier.EntityType) occupier.Occupier {
	for _, neighbor := range neighbors {
		if neighbor != nil && neighbor.Type() == target {
			return neighbor
		}
	}
	return nil
}
