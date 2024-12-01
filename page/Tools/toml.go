package tools

import (
	"fmt"
	"os"

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

func LoadPageSettings() (Page, error) {
	var config Config
	data, err := os.ReadFile("config/Settings.toml")
	if err != nil {
		return Page{}, fmt.Errorf("error reading file: %v", err)
	}
	if err := toml.Unmarshal(data, &config); err != nil {
		return Page{}, fmt.Errorf("error decoding TOML: %v", err)
	}
	return config.Page, nil
}
