package init_

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/spawner"
	"simulation/internal/worldmap"
)

type SpawnOccupiersAction struct {
	spawner.Spawner
}

func (s SpawnOccupiersAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	for entity := range occupier.TypeCount_ {
		for range s.AmountFor(entity, w.Size()) {
			summoned := s.New(entity, w.FreePos())
			w.Put(summoned.Pos(), summoned)
			o[entity] = append(o[entity], summoned.Pos())
		}
	}
}
