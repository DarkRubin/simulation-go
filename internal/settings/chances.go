package settings

import (
	"fmt"
	"simulation/internal/occupier"
)

func ChanceFor(entity occupier.EntityType, s Settings) uint {
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
		panic(fmt.Sprintf("unknown occupier type for spawn: %s", entity))
	}
}
