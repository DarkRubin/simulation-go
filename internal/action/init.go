package action

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/positon"
	"simulation/internal/worldmap"
)

type InitAction interface {
	Action
}

type PutEntitiesInitAction struct {
}

func (p PutEntitiesInitAction) Act(w *worldmap.Map) {
	for entType := range occupier.TypeCount_ {
		for range amountFor(entType, w.Size()) {
			pos := positon.RandomPos(w.Height(), w.Width())
			spawnOnMap(w, pos, entType)
		}
	}
}

func spawnOnMap(w *worldmap.Map, pos positon.Position, entType occupier.EntityType) {
	var occ occupier.Occupier
	if entType.IsCreature() {
		occ = creature.NewCreature(entType, pos, 10, 3)
	} else {
		occ = occupier.NewStaticOccupier(entType, pos)
	}
	w.SetOccupier(pos, occ)
}

func amountFor(entType occupier.EntityType, available uint) uint {
	chance := 20 / float64(100)
	return uint(float64(available) * chance)
}
