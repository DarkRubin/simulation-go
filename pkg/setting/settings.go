package setting

type Settings struct {
	HerbivoreSettings
	PredatorSettings
	SpawnChanceSettings
	WorldMapSettings
	MoveSettings
}

func Default() *Settings {
	return &Settings{
		DefaultHerbivoreSettings(),
		DefaultPredatorSettings(),
		DefaultSpawnChanceSettings(),
		DefaultWorldMapSettings(),
		DefaultMoveSettings(),
	}
}
