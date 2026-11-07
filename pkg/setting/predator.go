package setting

type PredatorSettings struct {
	PredatorSpeed      uint
	PredatorHp         uint
	PredatorAttackHeal int
	PredatorKillHeal   int
	PredatorDamage     uint
}

func DefaultPredatorSettings() PredatorSettings {
	return PredatorSettings{
		PredatorSpeed:      1,
		PredatorHp:         10,
		PredatorAttackHeal: 3,
		PredatorKillHeal:   5,
		PredatorDamage:     5,
	}
}
