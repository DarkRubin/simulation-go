package ui

type Dialog interface {
	Welcome()
	Goodbye()
	InvalidCommand()
}
