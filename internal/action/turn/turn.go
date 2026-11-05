package turn

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/worldmap"
	"slices"
)

func forEachCreature(w *worldmap.Map, o occupier.Occupiers, action func(creature creature.Creature)) {
	for _, creatureType := range creature.Types {
		for _, creatPos := range o[creatureType] {
			if slices.Index(o[creatureType], creatPos) == -1 {
				continue
			}
			occ := w.Get(creatPos)
			if occ == nil {
				continue
			}
			creat := occ.(creature.Creature)
			action(creat)
		}
	}
}
