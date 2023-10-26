package market_test

import (
	"net/url"
	"testing"

	"github.com/arensusu/pionex/domain"
	"github.com/arensusu/pionex/market"
	"github.com/arensusu/pionex/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetTrades(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	query := url.Values{}
	query.Add("Symbol", "TRB_USDT_PERP")
	query.Add("Limit", "1")

	data := `{
		"result": true,
		"data": {
			"trades": [
				{
					"symbol": "TRB_USDT_PERP",
					"tradeId": "200000000016879856",
					"price": "112.48",
					"size": "0.12",
					"side": "BUY",
					"timestamp": 1698330417889
				}
			]
		},
		"timestamp": 1698330418113
	}`
	mockClient := mocks.NewMockClient(ctrl)
	mockClient.EXPECT().HttpGet("/api/v1/market/trades", query).Return([]byte(data), nil)

	param := market.GetTradesParam{
		Symbol: "TRB_USDT_PERP",
		Limit:  func(val int) *int { return &val }(1),
	}

	service := market.NewMarketService(mockClient)
	res, err := service.GetTrades(param)

	expect := market.GetTradesResponse{
		HttpResponse: domain.HttpResponse{
			Result:    true,
			Timestamp: 1698330418113,
		},
		Data: market.GetTradesData{
			Trades: []market.GetTradesTrades{
				{
					Symbol:    "TRB_USDT_PERP",
					TradeId:   "200000000016879856",
					Price:     "112.48",
					Size:      "0.12",
					Side:      "BUY",
					Timestamp: 1698330417889,
				},
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expect, *res)
}
