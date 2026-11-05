package init_

import (
	"fmt"
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/pos"
	"simulation/internal/settings"
	"simulation/internal/worldmap"
)

type SpawnOccupiersAction struct {
	*settings.Settings
}

func (s SpawnOccupiersAction) Act(w *worldmap.Map, m occupier.Occupiers) {
	for entity := range occupier.TypeCount_ {
		amount := calculateAmount(w, s.chanceFor(entity))
		m[entity] = s.spawnAll(entity, w, amount)
	}
}

func (s SpawnOccupiersAction) spawnAll(entity occupier.EntityType, w *worldmap.Map, amount int) []pos.Position {
	positions := make([]pos.Position, 0, amount)
	for len(positions) != amount {
		position := pos.RandomPos(w.Width(), w.Height())
		if !w.IsPosFree(position) {
			continue
		}

		var occ occupier.Occupier
		switch entity {
		case occupier.TypeHerbivore:
			occ = creature.NewHerbivore(position, int(s.HerbivoreHp), s.HerbivoreSpeed)
		case occupier.TypePredator:
			occ = creature.NewPredator(position, s.PredatorDamage, int(s.PredatorHp), s.PredatorSpeed)
		default:
			occ = occupier.NewStaticOccupier(entity, position)
		}
		w.Put(position, occ)

		positions = append(positions, position)
	}
	return positions
}

func (s SpawnOccupiersAction) chanceFor(e occupier.EntityType) uint {
	switch e {
	case occupier.TypeGrass:
		return s.GrassSpawnChance
	case occupier.TypeRock:
		return s.RockSpawnChance
	case occupier.TypeTree:
		return s.TreeSpawnChance
	case occupier.TypeHerbivore:
		return s.HerbivoreSpawnChance
	case occupier.TypePredator:
		return s.PredatorSpawnChance
	default:
		panic(fmt.Sprintf("unknown occupier type for spawn: %s", e))
	}
}

func calculateAmount(w *worldmap.Map, chance uint) int {
	percent := float64(chance) / 100.0
	return int(float64(w.Size()) * percent)
}
