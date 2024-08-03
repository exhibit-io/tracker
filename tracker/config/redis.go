package config

import (
	"fmt"
	"net/url"
	"strconv"
)

// RedisConfig represents the configuration for Redis.
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// GetURL constructs and returns the Redis URL as a *url.URL type.
func (r RedisConfig) GetURL() (*url.URL, error) {
	url := &url.URL{
		Scheme: "redis",
		Host:   fmt.Sprintf("%s:%s", r.Host, r.Port),
		Path:   strconv.Itoa(r.DB),
	}

	return url, nil
}

// GetAddr returns the Redis address in the format host:port.
func (r RedisConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", r.Host, r.Port)
}

// LoadRedisConfig loads the Redis configuration from environment variables.
func LoadRedisConfig() RedisConfig {
	return RedisConfig{
		Host:     getEnv("REDIS_HOST", "localhost"),
		Port:     getEnv("REDIS_PORT", "6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       getEnvAsInt("REDIS_DB", 0),
	}
}
