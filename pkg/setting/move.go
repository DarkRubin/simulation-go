package setting

import "time"

type MoveSettings struct {
	MoveDelayMs uint
}

func (s *MoveSettings) GetMoveDelayDuration() time.Duration {
	return time.Duration(s.MoveDelayMs) * time.Millisecond
}

func DefaultMoveSettings() MoveSettings {
	return MoveSettings{
		MoveDelayMs: 600,
	}
}
