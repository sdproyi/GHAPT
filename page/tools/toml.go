package tools

import (
	"fmt"
	"os"
	"sync"

	toml "github.com/pelletier/go-toml/v2"
)

type Page struct {
	Logo        string `toml:"logo"`
	Title       string `toml:"title"`
	Description string `toml:"description"`
}

type Config struct {
	Page Page `toml:"page"`
}

var (
	pageSettings Page
	settingsOnce sync.Once
)

func PageSettings() Page {
	settingsOnce.Do(func() {
		data, err := os.ReadFile("config/Settings.toml")
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		var config Config
		if err := toml.Unmarshal(data, &config); err != nil {
			fmt.Printf("Error decoding TOML: %v\n", err)
			return
		}

		pageSettings = config.Page
	})
	return pageSettings
}
