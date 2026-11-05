package turn

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/spawner"
	"simulation/internal/worldmap"
)

type RespawnOccupiersAction struct {
	spawner.Spawner
}

func (r *RespawnOccupiersAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	for entity := range occupier.TypeCount_ {
		if len(o[entity]) > 0 {
			continue
		}
		for range r.AmountFor(entity, w.Size()) / 2 {
			summoned := r.New(entity, w.FreePos())
			w.Put(summoned.Pos(), summoned)
			o[entity] = append(o[entity], summoned.Pos())
		}
	}
}
