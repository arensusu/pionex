package account

import (
	"encoding/json"
	"net/url"

	"github.com/arensusu/pionex/domain"
)

type GetBalancesResponse struct {
	domain.HttpResponse
	Data GetBalancesData `json:"data"`
}

type GetBalancesData struct {
	Balances []GetBalancesBalance `json:"balances"`
}

type GetBalancesBalance struct {
	Coin   string `json:"coin"`
	Free   string `json:"free"`
	Frozen string `json:"frozen"`
}

func (s *AccountService) GetBalances() (*GetBalancesResponse, error) {

	body, err := s.client.HttpGet("/api/v1/account/balances", url.Values{})
	if err != nil {
		return nil, err
	}

	res := new(GetBalancesResponse)
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}
