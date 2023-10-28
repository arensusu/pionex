package market

import (
	"encoding/json"

	"github.com/arensusu/pionex/domain"
	"github.com/google/go-querystring/query"
)

type GetBookTickersParam struct {
	Symbol *string `json:"symbol"`
	Type   *string `json:"type"`
}

type GetBookTickersResponse struct {
	domain.HttpResponse
	Data GetBookTickersData `json:"data"`
}

type GetBookTickersData struct {
	Tickers []GetBookTickersTicker `json:"tickers"`
}

type GetBookTickersTicker struct {
	Symbol    string `json:"symbol"`
	Timestamp int64  `json:"timestamp"`
	BidPrice  string `json:"bidPrice"`
	BidSize   string `json:"bidSize"`
	AskPrice  string `json:"askPrice"`
	AskSize   string `json:"askSize"`
}

func (s *MarketService) GetBookTickers(req GetBookTickersParam) (*GetBookTickersResponse, error) {
	urlQuery, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	body, err := s.client.HttpGet("/api/v1/market/bookTickers", urlQuery)
	if err != nil {
		return nil, err
	}

	res := new(GetBookTickersResponse)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}
