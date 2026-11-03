package wayfinder

import (
	"simulation/internal/pos"
	"slices"
)

type Bfs struct {
}

func (b Bfs) Find(curr pos.Position, target pos.Position, check check) []pos.Position {
	way := findRecursive(curr, []pos.Position{target}, check)
	if len(way) > 1 {
		return way[:len(way)-1]
	}
	return []pos.Position{}
}

func findRecursive(src pos.Position, way []pos.Position, check check) []pos.Position {
	visited := []pos.Position{src}
	next := []pos.Position{src}
	target := way[len(way)-1]

	if slices.Contains(getNearPositions(target), src) {
		return way
	}

	for _, current := range next {
		for _, nearPos := range getNearPositions(current) {
			if !slices.Contains(visited, nearPos) && check(nearPos) {
				if nearPos == target {
					way = append(way, current)
					return findRecursive(src, way, check)
				}
				next = append(next, nearPos)
			}
		}
		visited = append(visited, current)
	}
	return []pos.Position{}
}

func getNearPositions(src pos.Position) []pos.Position {
	return []pos.Position{
		src.Up(),
		src.Down(),
		src.Left(),
		src.Right(),
	}
}
