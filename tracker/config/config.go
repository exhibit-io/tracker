package config

import "github.com/joho/godotenv"

// Config represents the main configuration struct containing all service configs.
type Config struct {
	Redis   RedisConfig
	Tracker TrackerConfig
	Cors    CorsConfig
}

// LoadConfig loads the configuration from environment variables.
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		return nil
	}
	return &Config{
		Redis:   LoadRedisConfig(),
		Tracker: LoadTrackerConfig(),
		Cors:    LoadCorsConfig(),
	}
}
