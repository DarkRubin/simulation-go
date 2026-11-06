package wayfinder

import (
	"simulation/internal/pos"
)

type WayFinder interface {
	Find(current pos.Position, targets []pos.Position, expand expand) ([]pos.Position, bool)
}

type expand func(position pos.Position) []pos.Position
