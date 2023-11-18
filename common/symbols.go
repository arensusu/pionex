package common

import (
	"encoding/json"

	"github.com/arensusu/pionex/constants"
	"github.com/arensusu/pionex/domain"
	"github.com/google/go-querystring/query"
)

type GetSymbolsParam struct {
	Symbol *string
	Type   *string
}

type GetSymbolsResponse struct {
	domain.HttpResponse
	Data GetSymbolsData `json:"data"`
}

type GetSymbolsData struct {
	Symbols []GetSymbolsSymbol `json:"symbols"`
}

type GetSymbolsSymbol struct {
	Symbol          string `json:"symbol"`
	Type            string `json:"type"`
	BaseCurrency    string `json:"baseCurrency"`
	QuoteCurrency   string `json:"quoteCurrency"`
	BasePrecision   int    `json:"basePrecision"`
	QuotePrecision  int    `json:"quotePrecision"`
	AmountPrecision int    `json:"amountPrecision"`
	MinAmount       string `json:"minAmount"`
	MinTradeSize    string `json:"minTradeSize"`
	MaxTradeSize    string `json:"maxTradeSize"`
	MinTradeDumping string `json:"minTradeDumping"`
	MaxTradeDumping string `json:"maxTradeDumping"`
	Enable          bool   `json:"enable"`
	BuyCeiling      string `json:"buyCeiling"`
	SellFloor       string `json:"sellFloor"`
}

func (s *CommonService) GetSymbols(param GetSymbolsParam) (*GetSymbolsResponse, error) {
	query, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	body, err := s.client.HttpGet(constants.COMMON_SYMBOLS, query)
	if err != nil {
		return nil, err
	}

	res := new(GetSymbolsResponse)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}
