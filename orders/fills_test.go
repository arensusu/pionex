package orders_test

import (
	"testing"

	"github.com/arensusu/pionex/constants"
	"github.com/arensusu/pionex/domain"
	"github.com/arensusu/pionex/mocks"
	"github.com/arensusu/pionex/orders"
	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetFills(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	param := orders.GetFillsParam{
		Symbol:    "BTC_USDT",
		StartTime: nil,
		EndTime:   nil,
	}

	query, _ := query.Values(param)

	data := `{
		"data": {
			"fills": [
				{
					"id": 1,
					"orderId": 1234567890,
					"symbol": "BTC_USDT",
					"side": "BUY",
					"role": "MAKER",
					"price": "30000.00",
					"size": "0.1000",
					"fee": "0.15",
					"feeCoin": "USDT",
					"timestamp": 1566676132311
				}
			]
		},
		"result": true,
		"timestamp": 1566691672311
	}`

	mockClient := mocks.NewMockClient(ctrl)
	mockClient.EXPECT().HttpGet(constants.ORDERS_FILLS, query).Times(1).Return([]byte(data), nil)

	s := orders.NewOrderService(mockClient)

	expectedResponse := &orders.GetFillsResponse{
		HttpResponse: domain.HttpResponse{
			Result:    true,
			Timestamp: 1566691672311,
		},
		Data: orders.GetFillsData{
			Fills: []orders.GetFillsFill{
				{
					Id:        1,
					OrderId:   1234567890,
					Symbol:    "BTC_USDT",
					Side:      "BUY",
					Role:      "MAKER",
					Price:     "30000.00",
					Size:      "0.1000",
					Fee:       "0.15",
					FeeCoin:   "USDT",
					Timestamp: 1566676132311,
				},
			},
		},
	}

	response, err := s.GetFills(param)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}
