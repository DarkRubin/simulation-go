package turn

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/worldmap"
)

type RecoverCreatureAction struct{}

func (r RecoverCreatureAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	forEachCreature(w, o, func(creature creature.Creature) {
		creature.Recover()
	})
}
