package config

import "fmt"

type TrackerConfig struct {
	Host         string
	Port         string
	CookieName   string
	CookieDomain string
}

func (t TrackerConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", t.Host, t.Port)
}

func LoadTrackerConfig() TrackerConfig {
	return TrackerConfig{
		Host:         getEnv("HOST", "localhost"),
		Port:         getEnv("PORT", "8080"),
		CookieName:   getEnv("COOKIE_NAME", "fingerprint"),
		CookieDomain: getEnv("COOKIE_DOMAIN", ""),
	}
}
