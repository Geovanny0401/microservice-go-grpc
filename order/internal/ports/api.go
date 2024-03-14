package ports

import "github.com/Geovanny0401/microservice-go-grpc/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
