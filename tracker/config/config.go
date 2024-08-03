package config

import "github.com/joho/godotenv"

// Config represents the main configuration struct containing all service configs.
type Config struct {
	Redis   RedisConfig
	Tracker TrackerConfig
}

// LoadConfig loads the configuration from environment variables.
func LoadConfig() *Config {
	godotenv.Load()
	return &Config{
		Redis:   LoadRedisConfig(),
		Tracker: LoadTrackerConfig(),
	}
}
