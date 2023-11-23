package orders_test

import (
	"testing"

	"github.com/arensusu/pionex/constants"
	"github.com/arensusu/pionex/domain"
	"github.com/arensusu/pionex/mocks"
	"github.com/arensusu/pionex/orders"
	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetOpenOrders(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	param := orders.GetOpenOrdersParam{
		Symbol: "BTC_USDT",
	}

	data := `{
		"data": {
		  "orders":[
			{
			  "orderId": 1234567890,
			  "symbol": "BTC_USDT",
			  "type": "LIMIT",
			  "side": "SELL",
			  "price": "30000.00",
			  "size": "0.1000",
			  "filledSize": "0.0500",
			  "filledAmount": "1500.00",
			  "fee":  "0.15",
			  "feeCoin":  "USDT",
			  "status": "OPEN",
			  "IOC": false,
			  "clientOrderId":  "9e3d93d6-e9a4-465a-a39c-2e48568fe194",
			  "source": "API",
			  "createTime": 1566676132311,
			  "updateTime": 1566676132311
			}
		  ]
		},
		"result": true,
		"timestamp": 1566676132311
	}`

	query, _ := query.Values(param)

	mockClient := mocks.NewMockClient(ctrl)
	mockClient.EXPECT().HttpGet(constants.ORDERS_OPEN_ORDERS, query).Return([]byte(data), nil)

	expect := orders.GetOpenOrdersResponse{
		HttpResponse: domain.HttpResponse{
			Result:    true,
			Timestamp: 1566676132311,
		},
		Data: orders.GetOpenOrdersData{
			Orders: []orders.GetOpenOrdersOrder{
				{
					OrderId:       1234567890,
					Symbol:        "BTC_USDT",
					Type:          "LIMIT",
					Side:          "SELL",
					Price:         "30000.00",
					Size:          "0.1000",
					FilledSize:    "0.0500",
					FilledAmount:  "1500.00",
					Fee:           "0.15",
					FeeCoin:       "USDT",
					Status:        "OPEN",
					IOC:           false,
					ClientOrderId: "9e3d93d6-e9a4-465a-a39c-2e48568fe194",
					Source:        "API",
					CreateTime:    1566676132311,
					UpdateTime:    1566676132311,
				},
			},
		},
	}

	service := orders.NewOrderService(mockClient)
	res, err := service.GetOpenOrders(param)

	require.NoError(t, err)
	require.Equal(t, expect, *res)
}
