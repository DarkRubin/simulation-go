package terminal

import "simulation/internal/comm"

type Key rune

const (
	KeyResume   Key = 'R'
	KeyPause        = 'P'
	KeyNextTurn     = 'N'
	KeyExit         = 'E'
)

var KeyCommand = map[Key]comm.Command{
	KeyResume:   comm.ComResume,
	KeyPause:    comm.ComPause,
	KeyNextTurn: comm.ComNextTurn,
	KeyExit:     comm.ComExit,
}

var CommandKey = map[comm.Command]Key{
	comm.ComResume:   KeyResume,
	comm.ComPause:    KeyPause,
	comm.ComNextTurn: KeyNextTurn,
	comm.ComExit:     KeyExit,
}
