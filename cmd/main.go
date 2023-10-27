package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/arensusu/pionex"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// c := pionex.NewClient("123", "NFqv4MB3hB0SOiEsJNDP9e0jDdKPWbDqS_Z1dbU4")

	// params := map[string]string{
	// 	"symbol": "BTC_USDT",
	// 	"limit":  "1",
	// }
	// result := c.Sign("GET", "/api/v1/trade/allOrders", "1655896754515", params, `{"symbol": "BTC_USDT"}`)
	// fmt.Println(result)

	c := pionex.NewClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

	param := map[string]string{
		"symbol":        "BTC_USDT_PERP",
		"side":          "BUY",
		"type":          "LIMIT",
		"clientOrderId": strconv.FormatInt(time.Now().UnixMilli(), 10),
		"price":         "120",
		"size":          "0.1",
	}

	resp, err := c.HttpPost("/api/v1/trade/order", param)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Printf("%s\n", resp)
	}
}
