package main

import (
	"fmt"
	"net/url"

	"github.com/arensusu/pionex"
)

func main() {
	// c := pionex.NewClient("123", "NFqv4MB3hB0SOiEsJNDP9e0jDdKPWbDqS_Z1dbU4")

	// params := map[string]string{
	// 	"symbol": "BTC_USDT",
	// 	"limit":  "1",
	// }
	// result := c.Sign("GET", "/api/v1/trade/allOrders", "1655896754515", params, `{"symbol": "BTC_USDT"}`)
	// fmt.Println(result)

	c := pionex.NewClient("6w8C8aN1vxQQv8iipA83e51SEZLEkMGdDrHJU6wXLd2mrL99KBk7vc7jaNxmHCYU6K", "S72JjcEb3DhoPsBXoNisNqMeAQ4kPMTtsnGmWRAPDauxHFJHHwpwBSdxsJsjAEzf")
	query := url.Values{}
	query.Add("symbol", "TRB_USDT_PERP")
	query.Add("limit", "10")
	resp, err := c.HttpGet("/api/v1/market/trades", query)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Printf("%s\n", resp)
	}
}
