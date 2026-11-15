package setting

import "simulation/internal/wayfinder"

type WayfinderSettings struct {
	Algorithm wayfinder.AlgorithmName `yaml:"algorithm"`
}

func DefaultWayfinderSettings() WayfinderSettings {
	return WayfinderSettings{
		Algorithm: "bfs",
	}
}
