package turn

import (
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/pos"
	"simulation/internal/wayfinder"
	"simulation/internal/worldmap"
)

type CreatureMoveAction struct {
	Wayfinder wayfinder.WayFinder
}

func (c CreatureMoveAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	forEachCreature(w, o, func(creature creature.Creature) {
		if creature.IsTired() || len(o[creature.TargetType()]) < 0 {
			return
		}
		targType := creature.TargetType()
		expandFunc := func(position pos.Position) []pos.Position {
			return freeNeighbors(position, targType, w)
		}
		way, found := c.Wayfinder.Find(creature.Pos(), o[targType], expandFunc)

		if !found || len(way) == 0 {
			return
		}

		oldPos := creature.Pos()
		newPos := way[len(way)-1]
		if uint(len(way)) > creature.Speed() {
			newPos = way[creature.Speed()]
		}
		creature.Move(newPos)
		w.Remove(oldPos)
		w.Put(newPos, creature)
		o.UpdatePos(creature.Type(), oldPos, newPos)
	})
}

func freeNeighbors(src pos.Position, targetType occupier.EntityType, w *worldmap.Map) []pos.Position {
	neighbors := src.Neighbors()
	result := make([]pos.Position, 0, len(neighbors))
	for _, neighbor := range neighbors {
		if w.IsValidPos(neighbor) && (w.IsPosFree(neighbor) || w.Get(neighbor).Type() == targetType) {
			result = append(result, neighbor)
		}
	}

	return result
}
