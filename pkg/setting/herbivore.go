package setting

type HerbivoreSettings struct {
	Speed   uint `yaml:"speed"`
	Hp      uint `yaml:"hp"`
	EatHeal int  `yaml:"eat_heal"`
}

func DefaultHerbivoreSettings() HerbivoreSettings {
	return HerbivoreSettings{
		Speed:   2,
		Hp:      10,
		EatHeal: 5,
	}
}
