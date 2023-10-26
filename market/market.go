package market

import "github.com/arensusu/pionex/domain"

type MarketService struct {
	client domain.Client
}

func NewMarketService(c domain.Client) *MarketService {
	return &MarketService{client: c}
}
