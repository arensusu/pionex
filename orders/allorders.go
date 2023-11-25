package orders

import (
	"encoding/json"

	"github.com/arensusu/pionex/constants"
	"github.com/arensusu/pionex/domain"
	"github.com/google/go-querystring/query"
)

type GetAllOrdersParam struct {
	Symbol    string `json:"symbol"`
	StartTime *int64 `json:"startTime"`
	EndTime   *int64 `json:"endTime"`
	Limit     *int   `json:"limit"`
}

type GetAllOrdersResponse struct {
	domain.HttpResponse
	Data GetAllOrdersData `json:"data"`
}

type GetAllOrdersData struct {
	Orders []GetAllOrdersOrder `json:"orders"`
}

type GetAllOrdersOrder struct {
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

func (s *OrderService) GetAllOrders(param GetAllOrdersParam) (*GetAllOrdersResponse, error) {

	query, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	body, err := s.client.HttpGet(constants.ORDERS_ALL_ORDERS, query)
	if err != nil {
		return nil, err
	}

	res := new(GetAllOrdersResponse)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

type CancelAllOrdersResponse struct {
	domain.HttpResponse
}

func (s *OrderService) CancelAllOrders(symbol string) (*CancelAllOrdersResponse, error) {
	body, err := s.client.HttpDelete(constants.ORDERS_ALL_ORDERS, map[string]string{"symbol": symbol})
	if err != nil {
		return nil, err
	}

	res := new(CancelAllOrdersResponse)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}
