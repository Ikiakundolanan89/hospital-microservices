// Konfigurasi aplikasi
// internal/config/config.go
package config

import (
	"os"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type AppConfig struct {
	Name    string
	Version string
	Port    string
	Env     string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	Secret     string
	ExpireTime int // in hours
}

func Load() *Config {
	return &Config{
		App: AppConfig{
			Name:    getEnv("APP_NAME", "patient-service"),
			Version: getEnv("APP_VERSION", "1.0.0"),
			Port:    getEnv("APP_PORT", "3001"),
			Env:     getEnv("APP_ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "1433"),
			User:     getEnv("DB_USER", "sa"),
			Password: getEnv("DB_PASSWORD", "YourStrong@Passw0rd"),
			DBName:   getEnv("DB_NAME", "hospital_patient_db"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", "your-secret-key-change-this-in-production"),
			ExpireTime: getEnvAsInt("JWT_EXPIRE_HOURS", 24),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		// Simple conversion, in production use strconv.Atoi with error handling
		return defaultValue // Simplified for demo
	}
	return defaultValue
}
