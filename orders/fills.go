package orders

import (
	"encoding/json"

	"github.com/arensusu/pionex/constants"
	"github.com/arensusu/pionex/domain"
	"github.com/google/go-querystring/query"
)

type GetFillsParam struct {
	Symbol    string `json:"symbol"`
	StartTime *int64 `json:"startTime"`
	EndTime   *int64 `json:"endTime"`
}

type GetFillsResponse struct {
	domain.HttpResponse
	Data GetFillsData `json:"data"`
}

type GetFillsData struct {
	Fills []GetFillsFill `json:"fills"`
}

type GetFillsFill struct {
	Id        int64  `json:"id"`
	OrderId   int64  `json:"orderId"`
	Symbol    string `json:"symbol"`
	Side      string `json:"side"`
	Role      string `json:"role"`
	Price     string `json:"price"`
	Size      string `json:"size"`
	Fee       string `json:"fee"`
	FeeCoin   string `json:"feeCoin"`
	Timestamp int64  `json:"timestamp"`
}

func (s *OrderService) GetFills(param GetFillsParam) (*GetFillsResponse, error) {

	query, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	body, err := s.client.HttpGet(constants.ORDERS_FILLS, query)
	if err != nil {
		return nil, err
	}

	res := new(GetFillsResponse)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}
