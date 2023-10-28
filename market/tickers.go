package market

import (
	"encoding/json"

	"github.com/arensusu/pionex/domain"
	"github.com/google/go-querystring/query"
)

type GetTickersParam struct {
	Symbol *string `json:"symbol"`
	Type   *string `json:"type"`
}

type GetTickersResponse struct {
	domain.HttpResponse
	Data GetTickersData `json:"data"`
}

type GetTickersData struct {
	Tickers []GetTickersTicker `json:"tickers"`
}

type GetTickersTicker struct {
	Symbol string `json:"symbol"`
	Time   int64  `json:"time"`
	Open   string `json:"open"`
	Close  string `json:"close"`
	Low    string `json:"low"`
	High   string `json:"high"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Count  int64  `json:"count"`
}

func (s *MarketService) GetTickers(req GetTickersParam) (*GetTickersResponse, error) {
	urlQuery, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	body, err := s.client.HttpGet("/api/v1/market/tickers", urlQuery)
	if err != nil {
		return nil, err
	}

	res := new(GetTickersResponse)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}
