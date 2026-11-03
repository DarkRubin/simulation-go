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

	for _, herb := range herbivores {
		herbivore := herb.(*creature.Herbivore)

		targetIdx, dist := nearestManhattan(herb.Pos(), grasses)
		target := grasses[targetIdx]
		if dist == 1 {
			w.RemoveOccupier(herb.Pos())
			herbivore.Eat(target)
			w.SetOccupier(target.Pos(), herb)

			o[occupier.TypeGrass] = slices.Delete(grasses, targetIdx, targetIdx+1)
			return
		}
		check := func(position pos.Position) bool {
			return w.IsValidPos(position) && w.IsPosFree(position)
		}

		way := m.Wayfinder.Find(herb.Pos(), target.Pos(), check)
		herbivore.Move(way)
	}
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
