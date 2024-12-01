package tools

import (
	"fmt"
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

type Settings struct {
	Logo        string `toml:"logo"`
	Title       string `toml:"title"`
	Description string `toml:"description"`
}

func LoadSettings() (Settings, error) {
	var settings Settings
	data, err := os.ReadFile("config/Settings.toml")
	if err != nil {
		return settings, fmt.Errorf("error reading file: %v", err)
	}
	if err := toml.Unmarshal(data, &settings); err != nil {
		return settings, fmt.Errorf("error decoding TOML: %v", err)
	}
	return settings, nil
}
