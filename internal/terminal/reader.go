package terminal

import (
	"bufio"
	"errors"
	"io"
	"os"
	"simulation/internal/comm"
	"strings"
	"unicode/utf8"
)

var ErrUnknownCommand = errors.New("unknown command provided")

type StdInListener struct {
	scanner *bufio.Scanner
}

func (t *StdInListener) AcceptCommand() (comm.Command, error) {
	for {
		command, err := t.nextCommand()
		if errors.Is(err, io.EOF) {
			return 0, err
		}
		if errors.Is(err, ErrUnknownCommand) {
			continue
		}
		if err != nil {
			return command, err
		}

		return command, nil
	}
}

func (t *StdInListener) nextCommand() (comm.Command, error) {
	if !t.scanner.Scan() {
		if err := t.scanner.Err(); err != nil {
			return 0, err
		}
		return 0, io.EOF
	}
	line := strings.TrimSpace(t.scanner.Text())
	if len(line) == 0 {
		return 0, errors.New("empty input")
	}

	if utf8.RuneCountInString(line) > 1 {
		return 0, errors.New("command too long")
	}

	r := rune(strings.ToUpper(line)[0])

	if command, ok := KeyCommand[Key(r)]; ok {
		return command, nil
	}
	return 0, ErrUnknownCommand
}

func NewStdInListener() comm.Listener {
	return &StdInListener{bufio.NewScanner(os.Stdin)}
}
