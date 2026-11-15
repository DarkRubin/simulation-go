package setting

import (
	"fmt"
	"simulation/internal/occupier"
)

// SpawnChanceSettings we should handle that sum of all chances
// are less than 100%, other vice return error, before spawn
// every thing on a map
type SpawnChanceSettings struct {
	Grass     uint `yaml:"grass"`
	Rock      uint `yaml:"rock"`
	Tree      uint `yaml:"tree"`
	Herbivore uint `yaml:"herbivore"`
	Predator  uint `yaml:"predator"`
}

func (s SpawnChanceSettings) ChanceFor(entity occupier.EntityType) uint {
	switch entity {
	case occupier.TypeGrass:
		return s.Grass
	case occupier.TypeRock:
		return s.Rock
	case occupier.TypeTree:
		return s.Tree
	case occupier.TypeHerbivore:
		return s.Herbivore
	case occupier.TypePredator:
		return s.Predator
	default:
		panic(fmt.Sprintf("unknown occupier type for summoner: %s", entity))
	}
}

func DefaultSpawnChanceSettings() SpawnChanceSettings {
	return SpawnChanceSettings{
		Grass:     10,
		Rock:      10,
		Tree:      10,
		Herbivore: 10,
		Predator:  10,
	}
}
