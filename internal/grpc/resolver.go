package grpc

import (
	"context"

	"github.com/arxdsilva/desafios-api/internal/service"
	log "github.com/sirupsen/logrus"
)

type Resolver struct {
	// api.Unimplemented...Server
	svc         service.Orders
	transformer OrdersTransformer
}

func NewResolver(svc service.Orders, tfmr OrdersTransformer) Resolver {
	return Resolver{
		svc:         svc,
		transformer: tfmr,
	}
}

func (r Resolver) GetOrder(ctx context.Context, apiReq interface{}) (apiTypeReturn interface{}, err error) {
	log.Info("[GetOrder]")
	return nil, nil
}
