package pionex

import (
	"github.com/arensusu/pionex/account"
	"github.com/arensusu/pionex/market"
	"github.com/arensusu/pionex/order"
)

func (c *Client) Market() *market.MarketService {
	return market.NewMarketService(c)
}

func (c *Client) Account() *account.AccountService {
	return account.NewAccountService(c)
}

func (c *Client) Order() *order.OrderService {
	return order.NewOrderService(c)
}
