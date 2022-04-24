package configs

type Configuration struct {
	Server      ServerConfig
	ProductsAPI ExternalServiceConfig
	PubSub      PubSubConfig
}

type ExternalServiceConfig struct {
	URL       string
	Key       string
	TimeoutMs int64
}
