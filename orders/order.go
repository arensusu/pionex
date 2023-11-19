package orders

import (
	"encoding/json"

	"github.com/arensusu/pionex/constants"
	"github.com/arensusu/pionex/domain"
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
