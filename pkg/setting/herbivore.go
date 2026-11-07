package setting

type HerbivoreSettings struct {
	HerbivoreSpeed   uint
	HerbivoreHp      uint
	HerbivoreEatHeal int
}

func DefaultHerbivoreSettings() HerbivoreSettings {
	return HerbivoreSettings{
		HerbivoreSpeed:   2,
		HerbivoreHp:      10,
		HerbivoreEatHeal: 5,
	}
}
