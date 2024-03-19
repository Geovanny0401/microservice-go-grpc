package ports

import (
	"context"

	"github.com/Geovanny0401/microservice-go-grpc/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
