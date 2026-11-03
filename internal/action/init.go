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
		amount := amountFor(entType, w.Size())
		m[entType] = p.spawnAll(entType, w, amount)
	}
}
func (p PutEntitiesInitAction) spawnAll(entType occupier.EntityType, w *worldmap.Map, amount int) []occupier.Occupier {
	occupiers := make([]occupier.Occupier, 0, amount)
	for len(occupiers) != amount {
		randPos := pos.RandomPos(w.Height(), w.Width())
		if !w.IsPosFree(randPos) {
			continue
		}
		var occ occupier.Occupier
		if entType.IsCreature() {
			occ = creature.NewCreature(entType, randPos, 10, 1)
		} else {
			occ = occupier.NewStaticOccupier(entType, randPos)
		}
		w.SetOccupier(randPos, occ)

		occupiers = append(occupiers, occ)
	}
	return occupiers
}

func amountFor(entType occupier.EntityType, available uint) int {
	chance := 10 / 100.0
	return int(float64(available) * chance)
}
