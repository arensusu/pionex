package orders_test

import (
	"testing"

	"github.com/arensusu/pionex/constants"
	"github.com/arensusu/pionex/domain"
	"github.com/arensusu/pionex/mocks"
	"github.com/arensusu/pionex/orders"
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
