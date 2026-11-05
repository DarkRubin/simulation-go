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
		if creature.Moved() || len(o[creature.TargetType()]) < 0 {
			return
		}
		targType := creature.TargetType()
		expandFunc := func(position pos.Position) []pos.Position {
			return freeNeighbors(position, targType, w)
		}
		way, found := c.Wayfinder.Find(creature.Pos(), o[targType], expandFunc)

		oldPos := creature.Pos()
		if found && creature.Move(way) {
			w.Remove(oldPos)
			w.Put(creature.Pos(), creature)
			o.UpdatePos(creature.Type(), oldPos, creature.Pos())
		}
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
