package orders

import (
	"encoding/json"

	"github.com/arensusu/pionex/constants"
	"github.com/arensusu/pionex/domain"
	"github.com/google/go-querystring/query"
)

type NewOrderParam struct {
	Symbol        string  `json:"symbol"`
	Side          string  `json:"side"`
	Type          string  `json:"type"`
	ClientOrderId *string `json:"clientOrderId"`
	Size          *string `json:"size"`
	Price         *string `json:"price"`
	Amount        *string `json:"amount"`
	IOC           *bool   `json:"IOC"`
}

type NewOrderResponse struct {
	domain.HttpResponse
	Data NewOrderData `json:"data"`
}

type NewOrderData struct {
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

func (s *OrderService) NewOrder(param NewOrderParam) (*NewOrderResponse, error) {
	body, err := s.client.HttpPost(constants.ORDERS_ORDER, param)
	if err != nil {
		return nil, err
	}

	res := new(NewOrderResponse)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

type GetOrderParam struct {
	Symbol  string `json:"symbol"`
	OrderId int64  `json:"orderId"`
}

type GetOrderResponse struct {
	domain.HttpResponse
	Data GetOrderData `json:"data"`
}

type GetOrderData struct {
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

func (s *OrderService) GetOrder(param GetOrderParam) (*GetOrderResponse, error) {
	query, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	body, err := s.client.HttpGet(constants.ORDERS_ORDER, query)
	if err != nil {
		return nil, err
	}

	res := new(GetOrderResponse)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

type CancelOrderParam struct {
	Symbol  string `json:"symbol"`
	OrderId int64  `json:"orderId"`
}

type CancelOrderResponse struct {
	domain.HttpResponse
}

func (s *OrderService) CancelOrder(param CancelOrderParam) (*CancelOrderResponse, error) {
	body, err := s.client.HttpDelete(constants.ORDERS_ORDER, param)
	if err != nil {
		return nil, err
	}

	res := new(CancelOrderResponse)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

/*

 */
