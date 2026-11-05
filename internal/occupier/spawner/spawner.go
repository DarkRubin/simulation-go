package spawner

import (
	"fmt"
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/pos"
	"simulation/pkg/setting"
)

type Spawner struct {
	setting.Settings
}

func (s Spawner) New(entity occupier.EntityType, pos pos.Position) occupier.Occupier {
	switch entity {
	case occupier.TypeHerbivore:
		return creature.NewHerbivore(pos, int(s.HerbivoreHp), s.HerbivoreSpeed)
	case occupier.TypePredator:
		return creature.NewPredator(pos, s.PredatorDamage, int(s.PredatorHp), s.PredatorSpeed)
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

func (s Spawner) AmountFor(entity occupier.EntityType, worldSize uint) uint {
	chance := s.chanceFor(entity)
	percent := float64(chance) / 100.0
	return uint(float64(worldSize) * percent)
}

func (s Spawner) chanceFor(entity occupier.EntityType) uint {
	switch entity {
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
		panic(fmt.Sprintf("unknown occupier type for summoner: %s", entity))
	}
}
