package market

import (
	"encoding/json"

	"github.com/arensusu/pionex/domain"
	"github.com/google/go-querystring/query"
)

type GetTradesParam struct {
	Symbol string
	Limit  *int
}

type GetTradesResponse struct {
	domain.HttpResponse
	Data GetTradesData `json:"data"`
}

type GetTradesData struct {
	Trades []GetTradesTrades `json:"trades"`
}

type GetTradesTrades struct {
	Symbol    string `json:"symbol"`
	TradeId   string `json:"tradeId"`
	Price     string `json:"price"`
	Size      string `json:"size"`
	Side      string `json:"side"`
	Timestamp int64  `json:"timestamp"`
}

func (s *MarketService) GetTrades(req GetTradesParam) (*GetTradesResponse, error) {
	urlQuery, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	body, err := s.client.HttpGet("/api/v1/market/trades", urlQuery)
	if err != nil {
		return nil, err
	}

	res := new(GetTradesResponse)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}
