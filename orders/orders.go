package orders

import "github.com/arensusu/pionex/domain"

type OrderService struct {
	client domain.Client
}

func NewOrderService(c domain.Client) *OrderService {
	return &OrderService{client: c}
}
