package main

import (
	"context"
	"net/http"
	"time"

	"calculator.com/configs"
	"calculator.com/infra/data"
	pubsubmessaging "calculator.com/infra/messaging/pubsub"
	fiberserver "calculator.com/infra/transports/http/fiber"
	"cloud.google.com/go/pubsub"

	"calculator.com/internal/application/services"
	"calculator.com/internal/engine"
	"calculator.com/pkg/httputils"
)

func main() {
	// Load configuration
	config := configs.Configuration{
		Server:      configs.NewServerConfigFromEnv(),
		ProductsAPI: configs.NewMeliApiConfigFromEnv(),
		PubSub:      configs.NewPubSubConfigFromEnv(),
	}

	// Create product service and dependencies
	productHttpClient := &http.Client{Timeout: time.Millisecond * time.Duration(config.ProductsAPI.TimeoutMs)}
	productApiClient := httputils.NewHTTPClient(config.ProductsAPI.URL, productHttpClient)

	pr := data.NewProductApiRepository(productApiClient)
	ps := services.NewProductService(pr)

	// Create voucher service and dependencies
	vr := data.NewVoucherRepository()
	vs := services.NewVoucherService(vr)

	// Create engine
	e := engine.NewServiceEngine(engine.ServiceRegistry{
		VoucherService: vs,
		ProductService: ps,
	})

	// pubsub client
	psc, err := pubsub.NewClient(context.Background(), config.PubSub.ProjectID)
	if err != nil {
		panic(err)
	}

	// Create messaging
	psi := pubsubmessaging.NewPubsubMessaging(psc)

	// Create server
	// server := ginserver.New(config.Server)
	server := fiberserver.New(config.Server)

	// Initialize server handlers
	server.InitializeDefaultHandlers(e)
	server.InitializeVoucherHandlers(e, psi)

	// Start server
	if err := server.Start(); err != nil {
		panic(err)
	}
}
