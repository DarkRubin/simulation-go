package wayfinder

import (
	"simulation/internal/pos"
	"slices"
)

const MaxWayTurns = 100

type Bfs struct {
}

func (b Bfs) Find(current pos.Position, targets []pos.Position, expand expand) ([]pos.Position, bool) {
	way := make([]pos.Position, 0, MaxWayTurns)

	target, success := bfs(current, targets, expand)
	way = append([]pos.Position{target}, way...)
	if !success {
		return way, false
	}

	for !slices.Contains(current.Neighbors(), target) {
		nextTargets := removeAll(expand(target), way)
		target, success = bfs(current, nextTargets, expand)
		if !success {
			return way, false
		}
		if len(way)+1 == MaxWayTurns {
			panic("too many way tiles found")
			return way, false
		}
		way = append([]pos.Position{target}, way...)
	}

	return way[:len(way)-1], true
}

func bfs(start pos.Position, targets []pos.Position, expand expand) (pos.Position, bool) {
	visited := make(map[pos.Position]bool)
	var current pos.Position
	queue := []pos.Position{start}

	for len(queue) > 0 {
		current, queue = queue[len(queue)-1], queue[:len(queue)-1]
		if slices.Contains(targets, current) {
			return current, true
		}
		visited[current] = true
		for _, child := range expand(current) {
			if visited[child] == false {
				queue = append(queue, child)
				visited[child] = true
			}
		}
	}
	return pos.Position{}, false
}

func removeAll(src []pos.Position, delete []pos.Position) []pos.Position {
	result := make([]pos.Position, 0, len(src))
	for _, srcPos := range src {
		if !slices.Contains(delete, srcPos) {
			result = append(result, srcPos)
		}
	}
	return result
}
