package config

import "os"

type Config struct {
	Port        string
	Environment string
	CORSOrigins []string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	corsOrigins := []string{"http://localhost:5173"} // SvelteKit default
	if env == "production" {
		corsOrigins = []string{"https://yourdomain.com"} // Update for production
	}

	return &Config{
		Port:        port,
		Environment: env,
		CORSOrigins: corsOrigins,
	}
}
