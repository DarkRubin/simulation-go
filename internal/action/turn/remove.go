package turn

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/worldmap"
)

type RemoveDeadCreaturesAction struct{}

func (r RemoveDeadCreaturesAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	forEachCreature(w, o, func(creature creature.Creature) {
		if creature.HP() > 1 {
			return
		}
		w.Remove(creature.Pos())
		o.RemovePos(creature.Type(), creature.Pos())
	})
}
