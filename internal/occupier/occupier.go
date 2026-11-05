package occupier

import (
	"simulation/internal/pos"
	"slices"
)

type Occupier interface {
	Type() EntityType
	Pos() pos.Position
}

type Occupiers map[EntityType][]pos.Position

func (o Occupiers) RemovePos(key EntityType, remove pos.Position) {
	ind := slices.Index(o[key], remove)
	if ind == -1 {
		panic("Occupier doesn't exist")
	}
	o[key] = slices.Delete(o[key], ind, ind+1)
}

func (o Occupiers) UpdatePos(key EntityType, old pos.Position, new pos.Position) {
	ind := slices.Index(o[key], old)
	if ind == -1 {
		panic("Occupier doesn't exist")
	}
	o[key][ind] = new
}
