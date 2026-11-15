package wayfinder

import (
	"simulation/internal/pos"
)

type AlgorithmName string

const (
	BreadthFirstSearch AlgorithmName = "bfs"
)

type WayFinder interface {
	Find(current pos.Position, targets []pos.Position, expand expand) ([]pos.Position, bool)
	Name() AlgorithmName
}

type expand func(position pos.Position) []pos.Position

func New(name AlgorithmName) WayFinder {
	switch name {
	case BreadthFirstSearch:
		return bfs{}
	default:
		return nil
	}
}
