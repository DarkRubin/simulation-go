package init_

import (
	"simulation/internal/action"
	"simulation/internal/occupier"
	"simulation/internal/settings"
	"simulation/internal/worldmap"
)

type SpawnOccupiersAction struct {
	settings.Settings
}

func (s SpawnOccupiersAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	for entity := range occupier.TypeCount_ {
		chance := settings.ChanceFor(entity, s.Settings)
		amount := action.CalculateAmount(w, chance)
		o[entity] = action.SpawnAll(entity, w, s.Settings, amount)
	}
}
