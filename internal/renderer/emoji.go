package renderer

import (
	"simulation/internal/occupier"
	"simulation/internal/pos"
	"simulation/internal/worldmap"
	"strings"
)

type EmojiRender struct {
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

func NewEmojiRenderer(bufSize int) Renderer {
	builder := strings.Builder{}
	builder.Grow(bufSize)

	return &EmojiRender{&builder}
}

func (e *EmojiRender) Render(m worldmap.Map) string {
	defer e.builder.Reset()
	for i := range m.Height() {
		for j := range m.Width() {
			cell := m.Get(pos.Pos(j, i))
			e.builder.WriteString(getEmoji(cell))
			e.builder.WriteRune(' ')
		}
		e.builder.WriteRune('\n')
	}
	e.builder.WriteRune('\n')

	return e.builder.String()
}

func getEmoji(occ occupier.Occupier) string {
	if occ != nil {
		return TypeEmoji[occ.Type()]
	}
	return EmptyCellEmoji
}
