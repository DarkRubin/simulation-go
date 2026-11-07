package turn

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/worldmap"
)

type HungerCreatureAction struct{}

func (h HungerCreatureAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	forEachCreature(w, o, func(creature creature.Creature) {
		creature.Hurt(1)
	})
}
