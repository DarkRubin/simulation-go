package turn

import (
	"fmt"
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/worldmap"
	"simulation/pkg/setting"
	"slices"
)

type PredatorsAttackAction struct {
	setting.Settings
}

func (p PredatorsAttackAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	for _, predPos := range o[occupier.TypePredator] {
		if slices.Index(o[occupier.TypePredator], predPos) == -1 {
			continue
		}

		predator, ok := w.Get(predPos).(*creature.Predator)
		if predator == nil || !ok {
			continue
		}

		neighbors := w.GetAll(predPos.Neighbors())
		occ := findNearby(neighbors, predator.TargetType())
		if occ == nil {
			continue
		}
		herbivore, ok := occ.(*creature.Herbivore)
		if !ok {
			panic(fmt.Sprintf("predator target is not a %s", occupier.TypeHerbivore))
		}

		predator.Attack(herbivore, p.PredatorAttackHeal)

		targetDead := herbivore.HP() < 1
		if targetDead {
			oldPos := predator.Pos()
			newPos := herbivore.Pos()

			predator.Eat(herbivore, p.PredatorKillHeal)

			w.Remove(oldPos)
			w.Put(newPos, predator)
			o.RemovePos(predator.TargetType(), newPos)
			o.UpdatePos(predator.Type(), oldPos, newPos)
		}
	}
}
