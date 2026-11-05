package renderer

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"simulation/internal/occupier"
	"simulation/internal/pos"
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

func NewCliRenderer(bufSize int) *CliRenderer {
	builder := strings.Builder{}
	builder.Grow(bufSize)

	return &CliRenderer{&builder}
}

func (c CliRenderer) Render(m worldmap.Map) {
	defer c.builder.Reset()

	for i := range m.Height() {
		for j := range m.Width() {
			posOccupier := m.Get(pos.Pos(j, i))
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

func (c CliRenderer) ClearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
