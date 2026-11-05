package turn

import (
	"simulation/internal/action"
	"simulation/internal/occupier"
	"simulation/internal/settings"
	"simulation/internal/worldmap"
)

type RespawnOccupiersAction struct {
	settings.Settings
}

func (r RespawnOccupiersAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	for entity := range occupier.TypeCount_ {
		if len(o[entity]) == 0 {
			chance := settings.ChanceFor(entity, r.Settings)
			amount := action.CalculateAmount(w, chance)
			o[entity] = action.SpawnAll(entity, w, r.Settings, amount/2)
		}
	}
}
