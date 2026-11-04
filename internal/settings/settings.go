package settings

import "time"

const (
	DefaultMoveMsDelay = 600
	DefaultWorldHeight = 6
	DefaultWorldWidth  = 7

	DefaultCreaturesMaxSpeed = 3
	DefaultCreaturesMinSpeed = 1
	DefaultHerbivoreHp       = 10
	DefaultHerbivoreHeal     = 5
	DefaultPredatorHp        = 10
	DefaultPredatorHeal      = 5
	DefaultPredatorDamage    = 5

	DefaultGrassSpawnChance     = 10
	DefaultRockSpawnChance      = 10
	DefaultTreeSpawnChance      = 10
	DefaultHerbivoreSpawnChance = 10
	DefaultPredatorSpawnChance  = 10
)

type Settings struct {
	MoveDelay            uint
	WorldHeight          uint
	WorldWidth           uint
	CreaturesMaxSpeed    uint
	CreaturesMinSpeed    uint
	HerbivoreHp          uint
	HerbivoreHeal        uint
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
		DefaultWorldHeight,
		DefaultWorldWidth,
		DefaultCreaturesMaxSpeed,
		DefaultCreaturesMinSpeed,
		DefaultPredatorDamage,
		DefaultHerbivoreHp,
		DefaultPredatorHp,
		DefaultHerbivoreHeal,
		DefaultPredatorHeal,
		DefaultGrassSpawnChance,
		DefaultRockSpawnChance,
		DefaultTreeSpawnChance,
		DefaultHerbivoreSpawnChance,
		DefaultPredatorSpawnChance,
	}
}
