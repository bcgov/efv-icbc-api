package config

import (
	"os"
	"time"
)

// Config holds application configuration values.
type Config struct {
	// Port the server listens on (default "8080").
	Port string

	// HTTP server timeouts
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// Load reads configuration from environment variables and returns a Config with sensible defaults.
func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Keep defaults consistent with previous main.go values
	return Config{
		Port:         port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}

// MustLoad is a convenience that returns a pointer to loaded Config.
func MustLoad() *Config {
	cfg := Load()
	return &cfg
}
