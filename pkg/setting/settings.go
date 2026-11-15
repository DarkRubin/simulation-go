package setting

type Settings struct {
	HerbivoreSettings   `yaml:"herbivore"`
	PredatorSettings    `yaml:"predator"`
	SpawnChanceSettings `yaml:"spawn_chance"`
	WorldMapSettings    `yaml:"world_map"`
	MoveSettings        `yaml:"move"`
	WayfinderSettings   `yaml:"wayfinder"`
}

func Default() *Settings {
	return &Settings{
		DefaultHerbivoreSettings(),
		DefaultPredatorSettings(),
		DefaultSpawnChanceSettings(),
		DefaultWorldMapSettings(),
		DefaultMoveSettings(),
		DefaultWayfinderSettings(),
	}
}
