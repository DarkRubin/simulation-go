package setting

type WorldMapSettings struct {
	Height uint `yaml:"height"`
	Width  uint `yaml:"width"`
}

func DefaultWorldMapSettings() WorldMapSettings {
	return WorldMapSettings{
		Height: 6,
		Width:  7,
	}
}
