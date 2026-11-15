package terminal

import (
	"fmt"
	"simulation/pkg/ui"
)

type StdOutDialog struct {
}

func (t *StdOutDialog) InvalidCommand() {
	fmt.Println("Invalid command provided")
}

func (t *StdOutDialog) Welcome() {
	fmt.Println("Welcome to Simulation!")
}

func (t *StdOutDialog) Goodbye() {
	fmt.Println("Goodbye!")
}

func NewStdOutDialog() ui.Dialog {
	return &StdOutDialog{}
}
