package turn

import (
	"simulation/internal/action"
	"simulation/internal/occupier"
	"simulation/internal/settings"
	"simulation/internal/worldmap"
)

type RespawnOccupiersAction struct {
	settings.Settings
	movesBeforeRespawn map[occupier.EntityType]uint
}

func (r *RespawnOccupiersAction) init() {
	r.movesBeforeRespawn = make(map[occupier.EntityType]uint)
	for entity := range occupier.TypeCount_ {
		r.movesBeforeRespawn[entity] = r.Settings.MovesBeforeRespawn
	}
}

func (r *RespawnOccupiersAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	if r.movesBeforeRespawn == nil {
		r.init()
	}
	for entity := range occupier.TypeCount_ {
		if len(o[entity]) > 0 {
			continue
		}
		r.movesBeforeRespawn[entity] -= 1
		if r.movesBeforeRespawn[entity] == 0 {
			chance := settings.ChanceFor(entity, r.Settings)
			amount := action.CalculateAmount(w, chance)
			o[entity] = action.SpawnAll(entity, w, r.Settings, amount/2)

			r.movesBeforeRespawn[entity] = r.Settings.MovesBeforeRespawn
		}
	}
}
