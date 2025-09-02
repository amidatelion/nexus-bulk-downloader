package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Config struct {
		APIKey      string `json:"apikey"`
		AutoExtract bool   `json:"autoextract"`
	} `json:"config"`
	Games map[string]map[string]string `json:"games"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// List available games from config
func (c *Config) GameOptions() []string {
	opts := []string{}
	for game := range c.Games {
		opts = append(opts, game)
	}
	return opts
}

// List available mods for a given game
func (c *Config) ModOptions(game string) []string {
	opts := []string{}
	for id, name := range c.Games[game] {
		opts = append(opts, fmt.Sprintf("%s - %s", id, name))
	}
	return opts
}

// Extract just the mod ID from a "1234 - Name" choice
func ExtractModID(choice string) string {
	return strings.Split(choice, " - ")[0]
}
