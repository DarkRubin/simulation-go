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

func (m *Map) Height() uint {
	return m.height
}

func (m *Map) Width() uint {
	return m.width
}

func (m *Map) Size() uint {
	return m.height * m.width
}

func (m *Map) GetOccupier(pos pos.Position) occupier.Occupier {
	return m.data[pos]
}

func (m *Map) SetOccupier(pos pos.Position, o occupier.Occupier) {
	m.data[pos] = o
}

func (m *Map) RemoveOccupier(pos pos.Position) {
	delete(m.data, pos)
}

func (m *Map) IsPosFree(pos positon.Position) bool {
	return m.data[pos] == nil
}
