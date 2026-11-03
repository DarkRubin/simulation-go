package action

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/pos"
	"simulation/internal/worldmap"
)

type InitAction interface {
	Action
}

type PutEntitiesInitAction struct {
}

func (p PutEntitiesInitAction) Act(w *worldmap.Map, m occupier.Occupiers) {
	for entType := range occupier.TypeCount_ {
		for range amountFor(entType, w.Size()) {
			randPos := pos.RandomPos(w.Height(), w.Width())
			occ := spawn(randPos, entType)
			w.SetOccupier(randPos, occ)
			m[entType] = append(m[entType], occ)
		}
	}
}

func spawn(pos pos.Position, entType occupier.EntityType) occupier.Occupier {
	if entType.IsCreature() {
		return creature.NewCreature(entType, pos, 10, 1)
	}

	return occupier.NewStaticOccupier(entType, pos)
}

func amountFor(entType occupier.EntityType, available uint) uint {
	chance := 20 / 100.0
	return uint(float64(available) * chance)
}
