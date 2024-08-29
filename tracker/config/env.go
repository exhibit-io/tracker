package config

import (
	"os"
	"strconv"
	"strings"
)

// getEnv retrieves an environment variable's value or returns a default value if not set.
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// getEnvAsInt retrieves an environment variable's value as an integer or returns a default value if not set.
func getEnvAsInt(key string, defaultValue int) int {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func getEnvAsArray(key string, defaultValue []string) []string {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	values := strings.Split(valueStr, ",")
	return values
}
