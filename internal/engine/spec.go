package engine

import "calculator.com/internal/application/services"

type ServiceEngine interface {
	GetVoucherService() services.VoucherService
	GetProductService() services.ProductService
}

type engine struct {
	vs services.VoucherService
	ps services.ProductService
}

var _ ServiceEngine = (*engine)(nil)

func (e engine) GetVoucherService() services.VoucherService {
	return e.vs
}

func (e engine) GetProductService() services.ProductService {
	return e.ps
}

type ServiceRegistry struct {
	VoucherService services.VoucherService
	ProductService services.ProductService
}

func NewServiceEngine(sr ServiceRegistry) ServiceEngine {
	return engine{
		vs: sr.VoucherService,
		ps: sr.ProductService,
	}
}
