package setting

import (
	"errors"
	"fmt"
	"simulation/internal/occupier"
)

type SpawnChanceSettings struct {
	grassSpawnChance     uint
	rockSpawnChance      uint
	treeSpawnChance      uint
	herbivoreSpawnChance uint
	predatorSpawnChance  uint
}

func SpawnChance(grass uint, rock uint, tree uint, herbivore uint, predator uint) (SpawnChanceSettings, error) {
	if (grass + rock + tree + herbivore + predator) > 100 {
		return SpawnChanceSettings{}, errors.New("total spawn chance of all entities must less than 100%")
	}
	return SpawnChanceSettings{
		grassSpawnChance:     grass,
		rockSpawnChance:      rock,
		treeSpawnChance:      tree,
		herbivoreSpawnChance: herbivore,
		predatorSpawnChance:  predator,
	}, nil
}

func (ss SpawnChanceSettings) ChanceFor(entity occupier.EntityType) uint {
	switch entity {
	case occupier.TypeGrass:
		return ss.grassSpawnChance
	case occupier.TypeRock:
		return ss.rockSpawnChance
	case occupier.TypeTree:
		return ss.treeSpawnChance
	case occupier.TypeHerbivore:
		return ss.herbivoreSpawnChance
	case occupier.TypePredator:
		return ss.predatorSpawnChance
	default:
		panic(fmt.Sprintf("unknown occupier type for summoner: %s", entity))
	}
}

func DefaultSpawnChanceSettings() SpawnChanceSettings {
	return SpawnChanceSettings{
		grassSpawnChance:     10,
		rockSpawnChance:      10,
		treeSpawnChance:      10,
		herbivoreSpawnChance: 10,
		predatorSpawnChance:  10,
	}
}
