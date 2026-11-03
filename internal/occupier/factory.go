package occupier

import (
	"fmt"
	"simulation/internal/pos"
)

func NewStaticOccupier(entType EntityType, pos pos.Position) Occupier {
	switch entType {
	case TypeGrass:
		return newGrass(pos)
	case TypeTree:
		return newTree(pos)
	case TypeRock:
		return newRock(pos)
	default:
		panic(fmt.Sprintf("no static occupier found for type %s", entType))
	}

}
