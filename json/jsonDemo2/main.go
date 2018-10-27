package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	dec "github.com/shopspring/decimal"
)

type ticker struct {
	TickerLetters string      `json:"symbol"`
	TickerPrice   dec.Decimal `json:"price,string"`
}

func main() {
	baseURL := "https://api.binance.com/api/v3/ticker/price?symbol=LTCBTC"
	resp, _ := http.Get(baseURL)

	defer resp.Body.Close()

	// Body is returned json object from call
	body, _ := ioutil.ReadAll(resp.Body)

	var tickerInfo ticker
	_ = json.Unmarshal(body, &tickerInfo)

	fmt.Println(tickerInfo.TickerLetters)
	fmt.Println(tickerInfo.TickerPrice)
}
