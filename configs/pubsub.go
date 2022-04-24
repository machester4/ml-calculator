package configs

import "os"

type PubSubConfig struct {
	ProjectID string
}

func NewPubSubConfigFromEnv() PubSubConfig {
	projectID := os.Getenv("PUBSUB_PROJECT_ID")
	if projectID == "" {
		panic("PUBSUB_PROJECT_ID is not set")
	}

	return PubSubConfig{
		ProjectID: projectID,
	}
}
