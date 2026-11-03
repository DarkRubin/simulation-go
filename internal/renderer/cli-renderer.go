package renderer

import (
	"fmt"
	"simulation/internal/occupier"
	"simulation/internal/positon"
	"simulation/internal/worldmap"
	"strings"
)

// CliRenderer store buf to
type CliRenderer struct {
	builder *strings.Builder
}

var TypeEmoji = map[occupier.EntityType]string{
	occupier.TypeGrass:     "🌾",
	occupier.TypeRock:      "🪨",
	occupier.TypeTree:      "🌲",
	occupier.TypeHerbivore: "🐰",
	occupier.TypePredator:  "🐺",
	occupier.TypeCount_:    "⛔️",
}

const EmptyCellEmoji = "🟩"

func NewCliRenderer(m *worldmap.Map) *CliRenderer {
	builder := strings.Builder{}
	capacity := int(m.Width()*m.Width()) * 2
	builder.Grow(capacity)

	return &CliRenderer{&builder}
}

func (c CliRenderer) Render(m worldmap.Map) {
	defer c.builder.Reset()

	for i := range m.Height() {
		for j := range m.Width() {
			pos := positon.Pos(i, j)
			posOccupier := m.GetOccupier(pos)
			if posOccupier != nil {
				c.builder.WriteString(TypeEmoji[posOccupier.Type()])
			} else {
				c.builder.WriteString(EmptyCellEmoji)
			}
			c.builder.WriteRune(' ')
		}
		c.builder.WriteRune('\n')
	}
	c.builder.WriteRune('\n')
	fmt.Println(c.builder.String())
}
