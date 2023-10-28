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

func TestGetTickers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	param := market.GetTickersParam{
		Symbol: func(val string) *string { return &val }("BTC_USDT_PERP"),
	}

	query, _ := query.Values(param)

	data := `{
		"result": true,
		"data": {
			"tickers": [
				{
					"symbol": "BTC_USDT_PERP",
					"time": 1698461980000,
					"open": "34093.8",
					"close": "34104.3",
					"low": "33404.0",
					"high": "34253.2",
					"volume": "7359.4347",
					"amount": "249995326.42332977",
					"count": 382358
				}
			]
		},
		"timestamp": 1698461980228
	}`
	mockClient := mocks.NewMockClient(ctrl)
	mockClient.EXPECT().HttpGet("/api/v1/market/tickers", query).Return([]byte(data), nil)

	service := market.NewMarketService(mockClient)
	res, err := service.GetTickers(param)

	expect := market.GetTickersResponse{
		HttpResponse: domain.HttpResponse{
			Result:    true,
			Timestamp: 1698461980228,
		},
		Data: market.GetTickersData{
			Tickers: []market.GetTickersTicker{
				{
					Symbol: "BTC_USDT_PERP",
					Time:   1698461980000,
					Open:   "34093.8",
					Close:  "34104.3",
					Low:    "33404.0",
					High:   "34253.2",
					Volume: "7359.4347",
					Amount: "249995326.42332977",
					Count:  382358,
				},
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expect, *res)
}
