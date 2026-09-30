package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Java     JavaConfig     `toml:"java"`
	Build    BuildConfig    `toml:"build"`
	Commands CommandsConfig `toml:"commands"`
}

type JavaConfig struct {
	Version string `toml:"version"`
}

type BuildConfig struct {
	Tool string `toml:"tool"`
}

type CommandsConfig struct {
	Dev   string `toml:"dev"`
	Test  string `toml:"test"`
	Build string `toml:"build"`
}

func Load() (Config, error) {
	var cfg Config

	data, err := os.ReadFile("jenvx.toml")
	if err != nil {
		return cfg, err
	}

	if err := toml.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
