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

func TestNewOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	body := orders.NewOrderParam{
		Symbol: "BTC_USDT",
		Side:   "BUY",
		Type:   "LIMIT",
	}
	data := `
	{ 
		"data": {
		  "orderId": 1234567890,
		  "clientOrderId":  "9e3d93d6-e9a4-465a-a39c-2e48568fe194"
		},
		"result": true,
		"timestamp": 1566691672311
	}`

	mockClient := mocks.NewMockClient(ctrl)
	mockClient.EXPECT().HttpPost(constants.ORDERS_ORDER, body).Times(1).Return([]byte(data), nil)

	expect := orders.NewOrderResponse{
		HttpResponse: domain.HttpResponse{
			Result:    true,
			Timestamp: 1566691672311,
		},
		Data: orders.NewOrderData{
			OrderId:       1234567890,
			ClientOrderId: "9e3d93d6-e9a4-465a-a39c-2e48568fe194",
		},
	}

	service := orders.NewOrderService(mockClient)
	res, err := service.NewOrder(body)

	require.NoError(t, err)
	require.Equal(t, expect, *res)
}

func TestGetOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	param := orders.GetOrderParam{
		Symbol:  "BTC_USDT",
		OrderId: 1234567890,
	}

	query, _ := query.Values(param)

	data := `{ 
		"data": {
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
		  "IOC":  false,
		  "clientOrderId":  "9e3d93d6-e9a4-465a-a39c-2e48568fe194",
		  "source": "API",
		  "createTime": 1566676132311,
		  "updateTime": 1566676132311
		},
		"result": true,
		"timestamp": 1566691672311
	}`

	mockClient := mocks.NewMockClient(ctrl)
	mockClient.EXPECT().HttpGet(constants.ORDERS_ORDER, query).Times(1).Return([]byte(data), nil)

	expect := orders.GetOrderResponse{
		HttpResponse: domain.HttpResponse{
			Result:    true,
			Timestamp: 1566691672311,
		},
		Data: orders.GetOrderData{
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
	}

	service := orders.NewOrderService(mockClient)
	res, err := service.GetOrder(param)

	require.NoError(t, err)
	require.Equal(t, expect, *res)
}

func TestCancelOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	param := orders.CancelOrderParam{
		Symbol:  "BTC_USDT",
		OrderId: 1234567890,
	}

	data := `{
		"result": true,
		"timestamp": 1566691672311
	}`

	mockClient := mocks.NewMockClient(ctrl)
	mockClient.EXPECT().HttpDelete(constants.ORDERS_ORDER, param).Times(1).Return([]byte(data), nil)

	expect := orders.CancelOrderResponse{
		HttpResponse: domain.HttpResponse{
			Result:    true,
			Timestamp: 1566691672311,
		},
	}

	service := orders.NewOrderService(mockClient)
	res, err := service.CancelOrder(param)

	require.NoError(t, err)
	require.Equal(t, expect, *res)
}
