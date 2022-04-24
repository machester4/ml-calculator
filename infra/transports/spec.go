package transports

import (
	"calculator.com/infra/messaging"
	"calculator.com/internal/engine"
)

type Transport interface {
	InitializeDefaultHandlers(engine.ServiceEngine)
	InitializeVoucherHandlers(engine.ServiceEngine, messaging.Messaging)
	Start() error
}
