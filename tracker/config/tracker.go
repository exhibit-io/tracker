package config

import "fmt"

type TrackerConfig struct {
	Host         string
	Port         string
	CookieConfig TrackerCookieConfig
}

type TrackerCookieConfig struct {
	Name     string
	Domain   string
	Secure   bool
	HttpOnly bool
}

func (t TrackerConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", t.Host, t.Port)
}

func LoadTrackerConfig() TrackerConfig {
	cookieConfig := TrackerCookieConfig{
		Name:     getEnv("COOKIE_NAME", "fingerprint"),
		Domain:   getEnv("COOKIE_DOMAIN", ""),
		Secure:   getEnvAsBool("COOKIE_SECURE", false),
		HttpOnly: getEnvAsBool("COOKIE_HTTPONLY", false),
	}
	return TrackerConfig{
		Host:         getEnv("HOST", "0.0.0.0"),
		Port:         getEnv("PORT", "8080"),
		CookieConfig: cookieConfig,
	}
}
