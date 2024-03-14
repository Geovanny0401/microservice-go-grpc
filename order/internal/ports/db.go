package ports

import "github.com/Geovanny0401/microservice-go-grpc/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
