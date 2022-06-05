package entity

import (
	"encoding/json"
	"fmt"
)

type Trade struct {
	ID     int     `json:"id"`
	Market int     `json:"market"`
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
	IsBuy  bool    `json:"is_buy"`
}

func NewTrade(text string) *Trade {
	var trade Trade
	err := json.Unmarshal([]byte(text), &trade)

	if err != nil {
		// TODO print error
		fmt.Println(err)
		return nil
	}

	return &trade
}
