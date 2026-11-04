package action

import (
	"fmt"
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/pos"
	"simulation/internal/settings"
	"simulation/internal/worldmap"
)

type InitAction interface {
	Action
}

type PutEntitiesInitAction struct {
	*settings.Settings
}

func (p PutEntitiesInitAction) Act(w *worldmap.Map, m occupier.Occupiers) {
	for entType := range occupier.TypeCount_ {
		percent := float64(p.chanceFor(entType)) / 100.0
		amount := int(float64(w.Size()) * percent)
		m[entType] = p.spawnAll(entType, w, amount)
	}
}
func (p PutEntitiesInitAction) spawnAll(entType occupier.EntityType, w *worldmap.Map, amount int) []occupier.Occupier {
	occupiers := make([]occupier.Occupier, 0, amount)
	for len(occupiers) != amount {
		randPos := pos.RandomPos(w.Height(), w.Width())
		if !w.IsPosFree(randPos) {
			continue
		}
		var occ occupier.Occupier
		if entType.IsCreature() {
			occ = creature.NewCreature(entType, randPos, 10, 1)
		} else {
			occ = occupier.NewStaticOccupier(entType, randPos)
		}
		w.SetOccupier(randPos, occ)

		occupiers = append(occupiers, occ)
	}
	return occupiers
}

func (p PutEntitiesInitAction) chanceFor(e occupier.EntityType) uint {
	switch e {
	case occupier.TypeGrass:
		return p.Settings.GrassSpawnChance
	case occupier.TypeRock:
		return p.Settings.RockSpawnChance
	case occupier.TypeTree:
		return p.Settings.TreeSpawnChance
	case occupier.TypeHerbivore:
		return p.Settings.HerbivoreSpawnChance
	case occupier.TypePredator:
		return p.Settings.PredatorSpawnChance
	default:
		panic(fmt.Sprintf("unknown occupier type for spawn: %s", e))
	}
}
