package worldmap

import (
	"errors"
	"simulation/internal/pos"
)

func (m *Map) FreePos() pos.Position {
	if !m.IsHaveMoreSpace() {
		panic(errors.New("no free space on map"))
	}
	for {
		position := pos.RandomPos(m.Width(), m.Height())
		if m.IsPosFree(position) {
			return position
		}
	}
}
