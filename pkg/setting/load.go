package setting

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

const configFilename = "config.yaml"

func Load() *Settings {
	config, err := os.ReadFile(configFilename)
	if err != nil {
		fmt.Printf("Failed to read config file: %s, use default settings instead\n", configFilename)
		return Default()
	}

	sets := &Settings{}
	err = yaml.Load(config, &sets)
	if err != nil {
		fmt.Printf("Failed to parse config file: %s, use default settings instead\n", configFilename)
		return Default()
	}
	return sets
}
