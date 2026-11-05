package action

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/pos"
	"simulation/internal/settings"
	"simulation/internal/worldmap"
)

func CalculateAmount(w *worldmap.Map, chance uint) int {
	percent := float64(chance) / 100.0
	return int(float64(w.Size()) * percent)
}

func SpawnAll(entity occupier.EntityType, w *worldmap.Map, s settings.Settings, amount int) []pos.Position {
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
