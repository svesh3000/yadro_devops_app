package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Version string
	Service string
	Author  string
	Port    string
	ApiKey  string `json:"-"`
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}

	return val
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("The .env file was not found: environment variables were used")
	}

	port := getEnv("PORT", "8000")

	err := validatePort(port)
	if err != nil {
		return nil, fmt.Errorf("invalid PORT: %w", err)
	}

	version := getEnv("VERSION", "1.0.1")
	author := getEnv("AUTHOR", "m.sveshnikov1")

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API_KEY not set")
	}

	return &Config{
		Version: version,
		Service: "weather",
		Author:  author,
		Port:    port,
		ApiKey:  apiKey,
	}, nil
}
