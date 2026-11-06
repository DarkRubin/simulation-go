package setting

type WorldMapSettings struct {
	WorldMapHeight uint
	WorldMapWidth  uint
}

func DefaultWorldMapSettings() WorldMapSettings {
	return WorldMapSettings{
		WorldMapHeight: 6,
		WorldMapWidth:  7,
	}
}
