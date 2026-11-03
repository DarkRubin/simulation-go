package wayfinder

import (
	"simulation/internal/pos"
)

type WayFinder interface {
	Find(curr pos.Position, target pos.Position, check check) []pos.Position
}

type check func(position pos.Position) (available bool)
