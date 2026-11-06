package setting

type HerbivoreSettings struct {
	HerbivoreSpeed uint
	HerbivoreHp    uint
	HerbivoreHeal  uint
}

func DefaultHerbivoreSettings() HerbivoreSettings {
	return HerbivoreSettings{
		HerbivoreSpeed: 2,
		HerbivoreHp:    10,
		HerbivoreHeal:  5,
	}
}
