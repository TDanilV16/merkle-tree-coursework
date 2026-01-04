package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App     AppConfig     `yaml:"app"`
	Logging LoggingConfig `yaml:"logging"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Debug   bool   `yaml:"debug"`
}

type LoggingConfig struct {
	Level     string `yaml:"level"`
	Format    string `yaml:"format"`
	AddSource bool   `yaml:"add_source"`
}

func Load(path string) (*Config, error) {
	var err error
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func LoadOrCreate(path string) *Config {
	config, err := Load(path)
	if err != nil {
		slog.Warn("Unable to load config", "error", err)
		return Default()
	}

	return config
}

func Default() *Config {
	return &Config{
		App: AppConfig{
			Name:    "unknown-project",
			Version: "0.0.1",
			Debug:   false,
		},
		Logging: LoggingConfig{
			Level:     "info",
			Format:    "text",
			AddSource: false,
		},
	}
}
