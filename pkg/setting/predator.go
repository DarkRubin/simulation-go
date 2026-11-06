package setting

type PredatorSettings struct {
	PredatorSpeed  uint
	PredatorHp     uint
	PredatorHeal   uint
	PredatorDamage uint
}

func DefaultPredatorSettings() PredatorSettings {
	return PredatorSettings{
		PredatorSpeed:  1,
		PredatorHp:     10,
		PredatorHeal:   5,
		PredatorDamage: 5,
	}
}
