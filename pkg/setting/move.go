package setting

import "time"

type MoveSettings struct {
	DelayMs uint `yaml:"delay_ms"`
}

func (s *MoveSettings) GetMoveDelayDuration() time.Duration {
	return time.Duration(s.DelayMs) * time.Millisecond
}

func DefaultMoveSettings() MoveSettings {
	return MoveSettings{
		DelayMs: 600,
	}
}
