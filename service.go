package pionex

import (
	"github.com/arensusu/pionex/account"
	"github.com/arensusu/pionex/market"
	"github.com/arensusu/pionex/orders"
)

func (c *Client) Market() *market.MarketService {
	return market.NewMarketService(c)
}

func (c *Client) Account() *account.AccountService {
	return account.NewAccountService(c)
}

func (c *Client) Order() *orders.OrderService {
	return orders.NewOrderService(c)
}
