package config

// CorsConfig represents the configuration for CORS.
type CorsConfig struct {
	AllowedOrigins   []string
	AllowCredentials bool
}

// LoadCorsConfig loads the CORS configuration from environment variables.
func LoadCorsConfig() CorsConfig {
	return CorsConfig{
		AllowedOrigins:   getEnvAsArray("CORS_ALLOWED_ORIGINS", []string{"*"}),
		AllowCredentials: getEnvAsBool("CORS_ALLOW_CREDENTIALS", false),
	}
}
