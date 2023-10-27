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

func TestGetDepth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	query := url.Values{}
	query.Add("Symbol", "TRB_USDT_PERP")
	query.Add("Limit", "1")

	data := `{ 
		"data": {
		  "bids": [
			  ["29658.37", "0.0123"]
		  ],
		  "asks": [
			  ["29658.47", "0.0345"]
		  ],
		  "updateTime": 1566676132311
		},
		"result": true,
		"timestamp": 1566691672311
	  }`
	mockClient := mocks.NewMockClient(ctrl)
	mockClient.EXPECT().HttpGet("/api/v1/market/depth", query).Return([]byte(data), nil)

	param := market.GetDepthParam{
		Symbol: "TRB_USDT_PERP",
		Limit:  func(val int) *int { return &val }(1),
	}

	service := market.NewMarketService(mockClient)
	res, err := service.GetDepth(param)

	expect := market.GetDepthResponse{
		HttpResponse: domain.HttpResponse{
			Result:    true,
			Timestamp: 1566691672311,
		},
		Data: market.GetDepthData{
			Bids: [][2]string{
				{
					"29658.37",
					"0.0123",
				},
			},
			Asks: [][2]string{
				{
					"29658.47",
					"0.0345",
				},
			},
			UpdateTime: 1566676132311,
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expect, *res)
}
