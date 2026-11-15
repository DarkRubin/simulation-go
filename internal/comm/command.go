package comm

const (
	ComPause Command = iota
	ComResume
	ComNextTurn
	ComExit
	ComCount_
)

type Command uint

var Description = map[Command]string{
	ComPause:    "pause",
	ComResume:   "resume",
	ComNextTurn: "next move",
	ComExit:     "exit",
}

type Listener interface {
	AcceptCommand() (Command, error)
}
