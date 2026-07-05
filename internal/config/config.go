package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	ExternalAPIURL    string
	HTTPClientTimeout time.Duration
	CacheTTL          time.Duration
}

func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	externalAPIURL := os.Getenv("EXTERNAL_API_URL")
	if externalAPIURL == "" {
		externalAPIURL = "https://jsonplaceholder.typicode.com"
	}

	httpClientTimeoutStr := os.Getenv("HTTP_CLIENT_TIMEOUT")
	if httpClientTimeoutStr == "" {
		httpClientTimeoutStr = "5s"
	}

	httpClientTimeout, err := time.ParseDuration(httpClientTimeoutStr)
	if err != nil {
		httpClientTimeout = 5 * time.Second
	}

	cacheTTLStr := os.Getenv("CACHE_TTL")
	if cacheTTLStr == "" {
		cacheTTLStr = "1m"
	}

	cacheTTL, err := time.ParseDuration(cacheTTLStr)
	if err != nil {
		cacheTTL = time.Minute
	}

	return Config{
		AppPort: appPort,

		ExternalAPIURL:    externalAPIURL,
		HTTPClientTimeout: httpClientTimeout,
		CacheTTL:          cacheTTL,
	}, nil
}
