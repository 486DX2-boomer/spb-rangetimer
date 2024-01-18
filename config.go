package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port     string `json:"port"`
	ServerIP string `json:"server_ip"`
}

func (c Config) Load() (Config, error) {
	configPath := "config.json"

	// Check if config file exists
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		// Config file doesn't exist, create a default one
		if err := createDefaultConfig(configPath); err != nil {
			return Config{}, fmt.Errorf("failed to create default config: %v", err)
		}
	}

	// Read and unmarshal config file
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config data: %v", err)
	}

	return config, nil
}

func createDefaultConfig(configPath string) error {
	// Create a default config struct with your desired default values
	defaultConfig := Config{
		Port:     ":8090",
		ServerIP: "127.0.0.1", // Default server IP
		// Add default values for additional configuration parameters
	}

	// Marshal the default config struct to JSON
	defaultConfigData, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal default config: %v", err)
	}

	// Write the default config to the config file
	if err := os.WriteFile(configPath, defaultConfigData, 0644); err != nil {
		return fmt.Errorf("failed to write default config to file: %v", err)
	}

	fmt.Println("Default config file created at", configPath)
	return nil
}
