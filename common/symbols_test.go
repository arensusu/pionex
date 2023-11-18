package common_test

import (
	"testing"

	"github.com/arensusu/pionex/common"
	"github.com/arensusu/pionex/constants"
	"github.com/arensusu/pionex/domain"
	"github.com/arensusu/pionex/mocks"
	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetSymbols(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	data := `
	{ 
		"data": {
		  "symbols":[
			{
			  "symbol": "BTC_USDT",
			  "type": "SPOT",
			  "baseCurrency": "BTC",
			  "quoteCurrency": "USDT",
			  "basePrecision": 6,
			  "quotePrecision": 2,
			  "amountPrecision": 8,
			  "minAmount": "10",
			  "minTradeSize": "0.000001",
			  "maxTradeSize": "1000",
			  "minTradeDumping": "0.000001",
			  "maxTradeDumping": "100",
			  "enable": true,
			  "buyCeiling": "1.1",
			  "sellFloor": "0.9"
			}
		  ]
		},
		"result": true,
		"timestamp": 1566676132311
	}`
	param := common.GetSymbolsParam{}
	query, _ := query.Values(param)

	mockClient := mocks.NewMockClient(ctrl)
	mockClient.EXPECT().HttpGet(constants.COMMON_SYMBOLS, query).Times(1).Return([]byte(data), nil)

	service := common.NewCommonService(mockClient)
	res, err := service.GetSymbols(param)

	expect := common.GetSymbolsResponse{
		HttpResponse: domain.HttpResponse{
			Result:    true,
			Timestamp: 1566676132311,
		},
		Data: common.GetSymbolsData{
			Symbols: []common.GetSymbolsSymbol{
				{
					Symbol:          "BTC_USDT",
					Type:            "SPOT",
					BaseCurrency:    "BTC",
					QuoteCurrency:   "USDT",
					BasePrecision:   6,
					QuotePrecision:  2,
					AmountPrecision: 8,
					MinAmount:       "10",
					MinTradeSize:    "0.000001",
					MaxTradeSize:    "1000",
					MinTradeDumping: "0.000001",
					MaxTradeDumping: "100",
					Enable:          true,
					BuyCeiling:      "1.1",
					SellFloor:       "0.9",
				},
			},
		},
	}

	require.NoError(t, err)
	require.Equal(t, expect, *res)
}
