package transports

import (
	"calculator.com/internal/engine"
)

type Transport interface {
	InitializeDefaultHandlers(engine.ServiceEngine)
	InitializeVoucherHandlers(engine.ServiceEngine)
}
