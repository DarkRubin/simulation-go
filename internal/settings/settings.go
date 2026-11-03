package settings

import "time"

const (
	DefaultMoveDelay   = 400 * time.Millisecond
	DefaultWorldHeight = 6
	DefaultWorldWidth  = 7

	DefaultCreaturesMaxSpeed = 3
	DefaultCreaturesMinSpeed = 1
	DefaultHerbivoreHp       = 10
	DefaultHerbivoreHeal     = 5
	DefaultPredatorHp        = 10
	DefaultPredatorHeal      = 5
	DefaultPredatorDamage    = 5
)

type Settings struct {
	moveDelay         time.Duration
	worldHeight       uint
	worldWidth        uint
	creaturesMaxSpeed uint
	creaturesMinSpeed uint
	herbivoreHp       uint
	herbivoreHeal     uint
	predatorHp        uint
	predatorHeal      uint
	predatorDamage    uint
}

func (s *Settings) MoveDelay() time.Duration {
	return s.moveDelay
}

func (s *Settings) WorldHeight() uint {
	return s.worldHeight
}

func (s *Settings) WorldWidth() uint {
	return s.worldWidth
}

func Default() *Settings {
	return &Settings{
		DefaultMoveDelay,
		DefaultWorldHeight,
		DefaultWorldWidth,
		DefaultCreaturesMaxSpeed,
		DefaultCreaturesMinSpeed,
		DefaultPredatorDamage,
		DefaultHerbivoreHp,
		DefaultPredatorHp,
		DefaultHerbivoreHeal,
		DefaultPredatorHeal,
	}
}
