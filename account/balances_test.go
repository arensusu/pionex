package account_test

import (
	"net/url"
	"testing"

	"github.com/arensusu/pionex/account"
	"github.com/arensusu/pionex/domain"
	"github.com/arensusu/pionex/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetBalances(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	data := `{
		"result": true,
		"data": {
			"balances": [
				{
					"coin": "USDT",
					"free": "699.9082361133759589202",
					"frozen": "0"
				}
			]
		},
		"timestamp": 1700295763766
	}`

	mockClient := mocks.NewMockClient(ctrl)
	mockClient.EXPECT().HttpGet("/api/v1/account/balances", url.Values{}).Times(1).Return([]byte(data), nil)

	expect := account.GetBalancesResponse{
		HttpResponse: domain.HttpResponse{
			Result:    true,
			Timestamp: 1700295763766,
		},
		Data: account.GetBalancesData{
			Balances: []account.GetBalancesBalance{
				{
					Coin:   "USDT",
					Free:   "699.9082361133759589202",
					Frozen: "0",
				},
			},
		},
	}

	service := account.NewAccountService(mockClient)
	res, err := service.GetBalances()

	require.NoError(t, err)
	require.Equal(t, expect, *res)
}
