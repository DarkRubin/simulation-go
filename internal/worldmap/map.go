package worldmap

import (
	"simulation/internal/occupier"
	"simulation/internal/positon"
)

type Map struct {
	data   map[positon.Position]occupier.Occupier
	height uint
	width  uint
}

func New(h uint, w uint) *Map {
	return &Map{
		data:   make(map[positon.Position]occupier.Occupier, h*w),
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

func (m *Map) GetOccupier(pos positon.Position) occupier.Occupier {
	return m.data[pos]
}

func (m *Map) SetOccupier(pos positon.Position, o occupier.Occupier) {
	m.data[pos] = o
}

func (m *Map) RemoveOccupier(pos positon.Position) {
	delete(m.data, pos)
}

func (m *Map) IsPosFree(pos positon.Position) bool {
	return m.data[pos] == nil
}
