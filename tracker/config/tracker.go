package config

import "fmt"

type TrackerConfig struct {
	Host string
	Port string
}

func (t TrackerConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", t.Host, t.Port)
}

func LoadTrackerConfig() TrackerConfig {
	return TrackerConfig{
		Host: getEnv("HOST", "localhost"),
		Port: getEnv("PORT", "3000"),
	}
}
