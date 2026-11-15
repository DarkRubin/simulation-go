package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"simulation/internal/comm"
	"simulation/internal/renderer"
	"simulation/internal/worldmap"
	"strings"
)

const Contrast = "\u001B[30;47m"
const Reset = "\u001B[0m"

type Printer struct {
	renderer renderer.Renderer
	builder  strings.Builder
}

func NewPrinter(renderer renderer.Renderer) *Printer {
	return &Printer{
		renderer: renderer,
		builder:  strings.Builder{},
	}
}

func (p *Printer) CurrState(worldMap worldmap.Map, turnNumber int, paused bool) {
	clearScreen()

	p.TurnCount(turnNumber)
	p.builder.WriteString(p.renderer.Render(worldMap))
	p.Commands(paused)

	fmt.Println(p.builder.String())
}

func (p *Printer) Commands(paused bool) {
	for command := range comm.ComCount_ {
		if command == comm.ComPause && paused {
			continue
		} else if command == comm.ComResume && !paused {
			continue
		}
		p.writeKeyWithDescription(CommandKey[command], comm.Description[command])
	}
}

func (p *Printer) TurnCount(i int) {
	if i < 1 {
		p.builder.WriteString("Welcome to Simulation!\n")
	} else {
		p.builder.WriteString(fmt.Sprintf("\nMove №%d\n", i))
	}
}

func (p *Printer) writeKeyWithDescription(key Key, description string) {
	p.builder.WriteString(Contrast)
	p.builder.WriteString(string(key))
	p.builder.WriteString(Reset)
	p.builder.WriteString(" - ")
	p.builder.WriteString(description)
	p.builder.WriteString(" ")
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
