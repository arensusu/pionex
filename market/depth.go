package market

import (
	"encoding/json"

	"github.com/arensusu/pionex/domain"
	"github.com/google/go-querystring/query"
)

type GetDepthParam struct {
	Symbol string
	Limit  *int
}

type GetDepthResponse struct {
	domain.HttpResponse
	Data GetDepthData `json:"data"`
}

type GetDepthData struct {
	Bids       [][2]string `json:"bids"`
	Asks       [][2]string `json:"asks"`
	UpdateTime int64       `json:"updateTime"`
}

func (s *MarketService) GetDepth(req GetDepthParam) (*GetDepthResponse, error) {
	urlQuery, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	body, err := s.client.HttpGet("/api/v1/market/depth", urlQuery)
	if err != nil {
		return nil, err
	}

	res := new(GetDepthResponse)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}
