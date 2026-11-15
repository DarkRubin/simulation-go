package setting

type PredatorSettings struct {
	Speed      uint `yaml:"speed"`
	Hp         uint `yaml:"hp"`
	AttackHeal int  `yaml:"attack_heal"`
	KillHeal   int  `yaml:"kill_heal"`
	Damage     uint `yaml:"damage"`
}

func DefaultPredatorSettings() PredatorSettings {
	return PredatorSettings{
		Speed:      1,
		Hp:         10,
		AttackHeal: 3,
		KillHeal:   5,
		Damage:     5,
	}
}
