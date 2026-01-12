package config

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Environment constants
const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

// Log level constants
const (
	LogLevelDebug = "debug"
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"
)

// Log format constants
const (
	LogFormatJSON = "json"
	LogFormatText = "text"
)

type Config struct {
	Port        string
	Environment string
	CORSOrigins []string
	LogLevel    string
	LogFormat   string
}

func Load() (*Config, error) {
	// Load .env file if it exists (ignore error if file doesn't exist on prod environments for example)
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("PORT environment variable is required")
	}
	portNum, err := strconv.Atoi(port)
	if err != nil || portNum < 1 || portNum > 65535 {
		return nil, errors.New("PORT must be a valid port number (1-65535)")
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		return nil, errors.New("ENVIRONMENT environment variable is required")
	}
	if env != EnvDevelopment && env != EnvProduction {
		return nil, errors.New("ENVIRONMENT must be either 'development' or 'production'")
	}

	corsOriginsEnv := os.Getenv("CORS_ORIGINS")
	if corsOriginsEnv == "" {
		return nil, errors.New("CORS_ORIGINS environment variable is required")
	}

	corsOrigins := make([]string, 0)
	for origin := range strings.SplitSeq(corsOriginsEnv, ",") {
		if trimmed := strings.TrimSpace(origin); trimmed != "" {
			corsOrigins = append(corsOrigins, trimmed)
		}
	}

	if len(corsOrigins) == 0 {
		return nil, errors.New("CORS_ORIGINS contains no valid origins")
	}

	// Log configuration with defaults
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		if env == EnvDevelopment {
			logLevel = LogLevelDebug
		} else {
			logLevel = LogLevelInfo
		}
	}

	logFormat := os.Getenv("LOG_FORMAT")
	if logFormat == "" {
		if env == EnvDevelopment {
			logFormat = LogFormatText
		} else {
			logFormat = LogFormatJSON
		}
	}

	return &Config{
		Port:        port,
		Environment: env,
		CORSOrigins: corsOrigins,
		LogLevel:    logLevel,
		LogFormat:   logFormat,
	}, nil
}
