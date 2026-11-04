package action

import (
	"math"
	"simulation/internal/occupier"
	"simulation/internal/occupier/creature"
	"simulation/internal/pos"
	"simulation/internal/wayfinder"
	"simulation/internal/worldmap"
	"slices"
)

type TurnAction interface {
	Action
}

type MoveOrEatAction struct {
	Wayfinder wayfinder.WayFinder
}

func (m MoveOrEatAction) Act(w *worldmap.Map, o occupier.Occupiers) {
	herbivores := o[occupier.TypeHerbivore]
	grasses := o[occupier.TypeGrass]

	if len(herbivores) == 0 {
		return
	}

	for _, herb := range herbivores {
		if len(grasses) == 0 {
			return
		}

		herbivore := herb.(*creature.Herbivore)
		targetIdx, dist := nearestManhattan(herb.Pos(), grasses)
		target := grasses[targetIdx]
		if dist == 1 {
			w.Remove(herb.Pos())
			herbivore.Eat(target)
			w.Put(target.Pos(), herb)

			grasses = slices.Delete(grasses, targetIdx, targetIdx+1)
			continue
		}

		way := m.Wayfinder.Find(herb.Pos(), target.Pos(), func(position pos.Position) bool {
			return w.IsValidPos(position) && w.IsPosFree(position)
		})

		w.Remove(herbivore.Pos())
		herbivore.Move(way)
		w.Put(herbivore.Pos(), herb)
	}
	o[occupier.TypeGrass] = grasses
}

func nearestManhattan(src pos.Position, variants []occupier.Occupier) (int, uint) {
	smallestDist := uint(math.MaxUint)
	var result int
	for i, variant := range variants {
		xDiff := math.Abs(float64(int(src.X()) - int(variant.Pos().X())))
		yDiff := math.Abs(float64(int(src.Y()) - int(variant.Pos().Y())))

		distance := uint(xDiff + yDiff)
		if distance < smallestDist {
			smallestDist = distance
			result = i
		}
	}
	return result, smallestDist
}
