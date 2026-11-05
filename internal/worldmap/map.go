package worldmap

import (
	"simulation/internal/occupier"
	"simulation/internal/pos"
)

type Map struct {
	data   map[pos.Position]occupier.Occupier
	height uint
	width  uint
}

func New(h uint, w uint) *Map {
	return &Map{
		data:   make(map[pos.Position]occupier.Occupier, h*w),
		height: h,
		width:  w,
	}
}

func (m *Map) IsValidPos(pos pos.Position) bool {
	return pos.X() <= m.width-1 && pos.Y() <= m.height-1
}

func (m *Map) IsPosFree(pos pos.Position) bool {
	return m.data[pos] == nil
}

func (m *Map) Height() uint {
	return m.height
}

func (m *Map) Width() uint {
	return m.width
}

func (m *Map) Size() uint {
	return m.height * m.width
}

func (m *Map) Get(pos pos.Position) occupier.Occupier {
	return m.data[pos]
}

func (m *Map) Put(pos pos.Position, o occupier.Occupier) {
	m.data[pos] = o
}

func (m *Map) Remove(pos pos.Position) {
	delete(m.data, pos)
}
