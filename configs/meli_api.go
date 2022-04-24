package configs

import "os"

func NewMeliApiConfigFromEnv() ExternalServiceConfig {
	url := "https://api.mercadolibre.com"
	if urlEnv := os.Getenv("MELI_API_URL"); urlEnv != "" {
		url = urlEnv
	}

	return ExternalServiceConfig{
		URL:       url,
		Key:       "",
		TimeoutMs: 1000,
	}
}
