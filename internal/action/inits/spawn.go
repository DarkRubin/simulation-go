package inits

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/spawner"
	"simulation/internal/worldmap"
	"simulation/pkg/setting"
)

type SpawnOccupiersAction struct {
	setting.Settings
}

func (s SpawnOccupiersAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	for entity := range occupier.TypeCount_ {
		for range spawner.AmountFor(entity, s.Settings) {
			summoned := spawner.NewOccupier(entity, w.FreePos(), s.Settings)
			w.Put(summoned.Pos(), summoned)
			o[entity] = append(o[entity], summoned.Pos())
		}
	}
}
