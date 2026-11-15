package turn

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/worldmap"
	"simulation/pkg/setting"
	"slices"
)

type HerbivoresEatAction struct {
	setting.Settings
}

func (c HerbivoresEatAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	for _, herbPos := range o[occupier.TypeHerbivore] {
		if slices.Index(o[occupier.TypeHerbivore], herbPos) == -1 {
			continue
		}
		herbivore, ok := w.Get(herbPos).(*creature.Herbivore)
		if herbivore == nil || !ok {
			continue
		}

		if herbivore.IsTired() {
			continue
		}

		neighbors := w.GetAll(herbivore.Pos().Neighbors())
		food := findNearby(neighbors, herbivore.TargetType())
		if food != nil {
			oldPos := herbivore.Pos()
			newPos := food.Pos()

			herbivore.Eat(food, c.EatHeal)

			w.Remove(oldPos)
			w.Put(newPos, herbivore)
			o.RemovePos(herbivore.TargetType(), newPos)
			o.UpdatePos(herbivore.Type(), oldPos, newPos)
		}
	}
}
