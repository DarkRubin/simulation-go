package spawner

import (
	"fmt"
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/pos"
	"simulation/pkg/setting"
)

func NewOccupier(entity occupier.EntityType, pos pos.Position, s setting.Settings) occupier.Occupier {
	switch entity {
	case occupier.TypeHerbivore:
		return creature.NewHerbivore(pos, int(s.HerbivoreSettings.Hp), s.HerbivoreSettings.Speed)
	case occupier.TypePredator:
		return creature.NewPredator(pos, s.Damage, int(s.PredatorSettings.Hp), s.PredatorSettings.Speed)
	case occupier.TypeGrass:
		return occupier.NewGrass(pos)
	case occupier.TypeTree:
		return occupier.NewTree(pos)
	case occupier.TypeRock:
		return occupier.NewRock(pos)
	default:
		panic(fmt.Sprintf("no static occupier found for type %s", entity))
	}
}

func AmountFor(entity occupier.EntityType, s setting.Settings) uint {
	worldSize := s.Height * s.Width
	chance := s.ChanceFor(entity)
	percent := float64(chance) / 100.0
	return uint(float64(worldSize) * percent)
}
