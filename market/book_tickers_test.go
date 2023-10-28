package market_test

import (
	"testing"

	"github.com/arensusu/pionex/domain"
	"github.com/arensusu/pionex/market"
	"github.com/arensusu/pionex/mocks"
	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetBookTickers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	param := market.GetBookTickersParam{
		Symbol: func(val string) *string { return &val }("BTC_USDT_PERP"),
	}

	query, _ := query.Values(param)

	data := `{
		"result": true,
		"data": {
			"tickers": [
				{
					"symbol": "BTC_USDT_PERP",
					"bidPrice": "34067.2",
					"bidSize": "5.258",
					"askPrice": "34067.3",
					"askSize": "7.44",
					"timestamp": 1698463409738
				}
			]
		},
		"timestamp": 1698463409738
	}`
	mockClient := mocks.NewMockClient(ctrl)
	mockClient.EXPECT().HttpGet("/api/v1/market/bookTickers", query).Return([]byte(data), nil)

	service := market.NewMarketService(mockClient)
	res, err := service.GetBookTickers(param)

	expect := market.GetBookTickersResponse{
		HttpResponse: domain.HttpResponse{
			Result:    true,
			Timestamp: 1698463409738,
		},
		Data: market.GetBookTickersData{
			Tickers: []market.GetBookTickersTicker{
				{
					Symbol:    "BTC_USDT_PERP",
					Timestamp: 1698463409738,
					BidPrice:  "34067.2",
					BidSize:   "5.258",
					AskPrice:  "34067.3",
					AskSize:   "7.44",
				},
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expect, *res)
}
