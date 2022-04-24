package main

import (
	"net/http"
	"time"

	"calculator.com/configs"
	"calculator.com/infra/data"
	fiberserver "calculator.com/infra/transports/http/fiber"

	"calculator.com/internal/application/services"
	"calculator.com/internal/engine"
	"calculator.com/pkg/httputils"
)

func main() {
	config := configs.Configuration{
		Server: configs.ServerConfig{
			Port: ":8080",
		},
		ProductsAPI: configs.ExternalServiceConfig{
			URL:       "https://api.mercadolibre.com",
			TimeoutMs: 1000,
		},
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

	// Create server
	// server := ginserver.New(config.Server)
	server := fiberserver.New(config.Server)

	// Initialize server handlers
	server.InitializeDefaultHandlers(e)
	server.InitializeVoucherHandlers(e)

	// Start server
	if err := server.Start(); err != nil {
		panic(err)
	}
}
