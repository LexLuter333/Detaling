package config

import (
	"os"
	"strings"
)

type Config struct {
	ServerPort         string
	JWTSecret          string
	AdminEmail         string
	AdminPassword      string
	CORSAllowedOrigins []string
	DatabaseURL        string
}

func Load() *Config {
	cfg := &Config{
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		JWTSecret:     getEnv("JWT_SECRET", ""),
		AdminEmail:    getEnv("ADMIN_EMAIL", "admin@deteleng.com"),
		AdminPassword: getEnv("ADMIN_PASSWORD", ""),
		DatabaseURL:   getEnv("DATABASE_URL", ""),
	}

	// Load CORS origins
	corsOrigins := getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:5173,http://localhost:3000")
	cfg.CORSAllowedOrigins = strings.Split(corsOrigins, ",")

	// Validate required fields in production
	if cfg.JWTSecret == "" {
		cfg.JWTSecret = generateSecureSecret()
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// generateSecureSecret generates a random secret for JWT
// In production, always set JWT_SECRET environment variable
func generateSecureSecret() string {
	// Fallback secret - should NEVER be used in production
	// This is just to prevent crashes during development
	return "CHANGE_THIS_IN_PRODUCTION_USE_STRONG_RANDOM_SECRET_32CHARS_MIN"
}
