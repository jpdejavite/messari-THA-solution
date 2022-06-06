package usecase

import (
	"sync"

	"github.com/jpdejavite/messari-THA-solution/internal/domain/entity"
)

type ProcessTrade interface {
	Execute(text string)
	GetTradesSummary() map[int]entity.MarketSummary
}

type ProcessTradeImpl struct {
	tradesMap      map[int]entity.MarketSummary
	tradesMapMutex sync.Mutex
	finishProcess  sync.WaitGroup
}

func NewProcessTrade() ProcessTrade {
	return &ProcessTradeImpl{
		tradesMap:      make(map[int]entity.MarketSummary),
		tradesMapMutex: sync.Mutex{},
		finishProcess:  sync.WaitGroup{},
	}
}

func (pt *ProcessTradeImpl) Execute(text string) {
	pt.finishProcess.Add(1)
	go func(inputText string) {
		defer pt.finishProcess.Done()
		trade := entity.NewTrade(text)
		if trade == nil {
			return
		}

		pt.tradesMapMutex.Lock()
		defer pt.tradesMapMutex.Unlock()
		marketSummary, exists := pt.tradesMap[trade.Market]
		if !exists {
			marketSummary = entity.NewMarketSummary(*trade)
		}

		marketSummary.AddTrade(*trade)

		pt.tradesMap[trade.Market] = marketSummary
	}(text)
}

func (pt *ProcessTradeImpl) GetTradesSummary() map[int]entity.MarketSummary {
	pt.finishProcess.Wait()
	return pt.tradesMap
}
