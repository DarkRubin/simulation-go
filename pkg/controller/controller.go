package controller

import (
	"errors"
	"io"
	"simulation/internal/comm"
	"simulation/pkg/simulation"
	"simulation/pkg/ui"
)

func ListenAndManage(listener comm.Listener, d ui.Dialog, sim simulation.Simulation) {
	go sim.Start()

	for {
		command, err := listener.AcceptCommand()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			d.InvalidCommand()
			continue
		}
		switch command {
		case comm.ComPause:
			sim.Pause()
		case comm.ComResume:
			sim.Resume()
		case comm.ComNextTurn:
			sim.NextTurn()
		case comm.ComExit:
			d.Goodbye()
			return
		default:
			panic("Unknown command")
		}
	}
}
