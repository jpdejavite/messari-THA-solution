package usecase

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/jpdejavite/messari-THA-solution/internal/domain/entity"
)

type ProcessTrade interface {
	Execute(text string)
	GetTradesSummary() map[int]entity.MarketSummary
}

type ProcessTradeImpl struct {
	tradesMap map[int]entity.MarketSummary
}

func NewProcessTrade() ProcessTrade {
	return &ProcessTradeImpl{
		tradesMap: make(map[int]entity.MarketSummary),
	}
}

var tradesMap = make(map[int]entity.MarketSummary)
var tradesChannel = make(map[int]chan entity.Trade)
var mutex = sync.RWMutex{}

func (pt ProcessTradeImpl) Execute(text string) {
	trade := entity.NewTrade(text)
	if trade == nil {
		return
	}

	marketSummary, exists := pt.tradesMap[trade.Market]

	if !exists {
		marketSummary = entity.NewMarketSummary(*trade)
	}

	marketSummary.AddTrade(*trade)

	pt.tradesMap[trade.Market] = marketSummary
}

func (pt ProcessTradeImpl) GetTradesSummary() map[int]entity.MarketSummary {
	return pt.tradesMap
}

func processTrade(text string) {
	var trade entity.Trade
	err := json.Unmarshal([]byte(text), &trade)
	// _, err := simdjson.Parse([]byte(text), nil)

	if err != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}

	mutex.Lock()
	marketSummary, exists := tradesMap[trade.Market]

	if !exists {
		marketSummary = entity.MarketSummary{}
		marketSummary.Market = trade.Market
	}

	marketSummary.TotalVolume += trade.Volume
	marketSummary.MeanPrice += (trade.Price + float64(marketSummary.Count)*marketSummary.MeanPrice) / float64(marketSummary.Count+1)
	marketSummary.MeanVolume += (trade.Volume + float64(marketSummary.Count)*marketSummary.MeanVolume) / float64(marketSummary.Count+1)
	marketSummary.VolumeWeightedAveragePrice += (trade.Price/trade.Volume + float64(marketSummary.Count)*marketSummary.VolumeWeightedAveragePrice) / float64(marketSummary.Count+1)
	if trade.IsBuy {
		marketSummary.PercentageBuy += (1 + float64(marketSummary.Count)*marketSummary.PercentageBuy) / float64(marketSummary.Count+1)
	} else {
		marketSummary.PercentageBuy += float64(marketSummary.Count) * marketSummary.PercentageBuy / float64(marketSummary.Count+1)
	}
	marketSummary.Count++

	tradesMap[trade.Market] = marketSummary
	mutex.Unlock()

}
