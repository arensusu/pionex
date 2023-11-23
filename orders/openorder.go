package orders

import (
	"encoding/json"

	"github.com/arensusu/pionex/constants"
	"github.com/arensusu/pionex/domain"
	"github.com/google/go-querystring/query"
)

type GetOpenOrdersParam struct {
	Symbol string `json:"symbol"`
}

type GetOpenOrdersResponse struct {
	domain.HttpResponse
	Data GetOpenOrdersData `json:"data"`
}

type GetOpenOrdersData struct {
	Orders []GetOpenOrdersOrder `json:"orders"`
}

type GetOpenOrdersOrder struct {
	OrderId       int64  `json:"orderId"`
	Symbol        string `json:"symbol"`
	Type          string `json:"type"`
	Side          string `json:"side"`
	Price         string `json:"price"`
	Size          string `json:"size"`
	FilledSize    string `json:"filledSize"`
	FilledAmount  string `json:"filledAmount"`
	Fee           string `json:"fee"`
	FeeCoin       string `json:"feeCoin"`
	Status        string `json:"status"`
	IOC           bool   `json:"IOC"`
	ClientOrderId string `json:"clientOrderId"`
	Source        string `json:"source"`
	CreateTime    int64  `json:"createTime"`
	UpdateTime    int64  `json:"updateTime"`
}

func (s *OrderService) GetOpenOrders(param GetOpenOrdersParam) (*GetOpenOrdersResponse, error) {

	query, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	body, err := s.client.HttpGet(constants.ORDERS_OPEN_ORDERS, query)
	if err != nil {
		return nil, err
	}

	res := new(GetOpenOrdersResponse)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}
