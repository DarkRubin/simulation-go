package renderer

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"simulation/internal/occupier"
	"simulation/internal/pos"
	"simulation/internal/worldmap"
	"simulation/pkg/command"
	"strings"
	"unicode"
)

const Contrast = "\u001B[30;47m"
const Background = "\u001B[30;49m"
const Reset = "\u001B[0m"

type TuiRenderer struct {
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

func NewTuiRenderer(bufSize int) Renderer {
	builder := strings.Builder{}
	builder.Grow(bufSize)

	return &TuiRenderer{&builder}
}

func (c *TuiRenderer) Render(m worldmap.Map, move int, paused bool) {
	defer c.builder.Reset()
	clearScreen()

	c.TurnCount(move)
	c.WorldState(m)
	c.Commands(paused)

	fmt.Println(c.builder.String())
}

func (c *TuiRenderer) Commands(paused bool) {
	if paused {
		c.builder.WriteString("Input ")
		c.builder.WriteString(Contrast)
		c.builder.WriteRune(unicode.ToUpper(command.Key[command.SimResume]))
		c.builder.WriteString(Reset)
		c.builder.WriteString(" to start or ")
		c.builder.WriteString(Contrast)
		c.builder.WriteRune(unicode.ToUpper(command.Key[command.SimExit]))
		c.builder.WriteString(Reset)
		c.builder.WriteString(" to exit")
	} else {
		for comm := range command.SimUnknown {
			c.builder.WriteString(Contrast)
			c.builder.WriteRune(unicode.ToUpper(command.Key[comm]))
			c.builder.WriteString(Reset)
			c.builder.WriteString(" - ")
			c.builder.WriteString(command.Description[comm])
			c.builder.WriteString(" ")
		}
	}
}

func (c *TuiRenderer) TurnCount(i int) {
	if i < 1 {
		c.builder.WriteString("Welcome to Simulation!\n")
	} else {
		c.builder.WriteString(fmt.Sprintf("\nMove №%d\n", i))
	}
}

func (c *TuiRenderer) WorldState(m worldmap.Map) {
	c.builder.WriteString(Background)
	for i := range m.Height() {
		for j := range m.Width() {
			cell := m.Get(pos.Pos(j, i))
			c.builder.WriteString(getEmoji(cell))
			c.builder.WriteRune(' ')
		}
		c.builder.WriteRune('\n')
	}
	c.builder.WriteRune('\n')
	c.builder.WriteString(Reset)
}

func InvalidCommand() {
	fmt.Println("Invalid command provided")
}

func clearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func getEmoji(occ occupier.Occupier) string {
	if occ != nil {
		return TypeEmoji[occ.Type()]
	}
	return EmptyCellEmoji
}
