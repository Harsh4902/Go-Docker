package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var ConfigPath = filepath.Join(os.Getenv("HOME"), ".microcks-cli", "config.yaml")

type Config struct {
	Instance struct {
		Name        string `yaml:"name"`
		Image       string `yaml:"image"`
		Status      string `yaml:"status"`
		Port        string `yaml:"port"`
		ContainerID string `yaml:"containerID"`
		AutoRemove  bool   `yaml:"autoRemove"`
		Driver      string `yaml:"driver"`
	} `yaml:"instance"`
}

// Functions related to configs
func defaultConfig() *Config {
	return &Config{
		Instance: struct {
			Name        string `yaml:"name"`
			Image       string `yaml:"image"`
			Status      string `yaml:"status"`
			Port        string `yaml:"port"`
			ContainerID string `yaml:"containerID"`
			AutoRemove  bool   `yaml:"autoRemove"`
			Driver      string `yaml:"driver"`
		}{
			Name:       "microcks",
			Image:      "",
			Status:     "",
			Port:       "",
			AutoRemove: false,
			Driver:     "",
		},
	}
}

func EnsureConfig(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Config not found. Initializing default config.")
		cfg := defaultConfig()
		err := SaveConfig(path, cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to create default config: %w", err)
		}
		return cfg, nil
	}
	return LoadConfig(path)
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func SaveConfig(path string, cfg *Config) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
