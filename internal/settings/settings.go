package settings

import "time"

const (
	DefaultMoveMsDelay        = 600
	DefaultMovesBeforeRespawn = 3
	DefaultWorldHeight        = 6
	DefaultWorldWidth         = 7

	DefaultHerbivoreSpeed = 1
	DefaultHerbivoreHp    = 10
	DefaultHerbivoreHeal  = 5
	DefaultPredatorSpeed  = 1
	DefaultPredatorHp     = 10
	DefaultPredatorHeal   = 5
	DefaultPredatorDamage = 5

	DefaultGrassSpawnChance     = 10
	DefaultRockSpawnChance      = 10
	DefaultTreeSpawnChance      = 10
	DefaultHerbivoreSpawnChance = 10
	DefaultPredatorSpawnChance  = 10
)

type Settings struct {
	MoveDelay            uint
	MovesBeforeRespawn   uint
	WorldHeight          uint
	WorldWidth           uint
	HerbivoreSpeed       uint
	HerbivoreHp          uint
	HerbivoreHeal        uint
	PredatorSpeed        uint
	PredatorHp           uint
	PredatorHeal         uint
	PredatorDamage       uint
	GrassSpawnChance     uint
	RockSpawnChance      uint
	TreeSpawnChance      uint
	HerbivoreSpawnChance uint
	PredatorSpawnChance  uint
}

func (s *Settings) GetMoveDelayDuration() time.Duration {
	return time.Duration(DefaultMoveMsDelay) * time.Millisecond
}

func Default() *Settings {
	return &Settings{
		DefaultMoveMsDelay,
		DefaultMovesBeforeRespawn,
		DefaultWorldHeight,
		DefaultWorldWidth,
		DefaultHerbivoreSpeed,
		DefaultHerbivoreHp,
		DefaultHerbivoreHeal,
		DefaultPredatorSpeed,
		DefaultPredatorHp,
		DefaultPredatorHeal,
		DefaultPredatorDamage,
		DefaultGrassSpawnChance,
		DefaultRockSpawnChance,
		DefaultTreeSpawnChance,
		DefaultHerbivoreSpawnChance,
		DefaultPredatorSpawnChance,
	}
}
