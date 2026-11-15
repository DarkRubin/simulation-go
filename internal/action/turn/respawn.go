package turn

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/spawner"
	"simulation/internal/worldmap"
	"simulation/pkg/setting"
)

type RespawnOccupiersAction struct {
	setting.Settings
}

func (r RespawnOccupiersAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	for entity := range occupier.TypeCount_ {
		if len(o[entity]) > 0 {
			continue
		}
		for range spawner.AmountFor(entity, r.Settings) / 2 {
			summoned := spawner.NewOccupier(entity, w.FreePos(), r.Settings)
			w.Put(summoned.Pos(), summoned)
			o[entity] = append(o[entity], summoned.Pos())
		}
	}
}
