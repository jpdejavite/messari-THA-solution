package entity

type MarketSummary struct {
	Market                     int     `json:"market"`
	TotalVolume                float64 `json:"total_volume"`
	MeanPrice                  float64 `json:"mean_price"`
	MeanVolume                 float64 `json:"mean_volume"`
	VolumeWeightedAveragePrice float64 `json:"volume_weighted_average_price"`
	PercentageBuy              float64 `json:"percentage_buy"`
	Count                      int     `json:"-"`
}

func NewMarketSummary(trade Trade) MarketSummary {
	return MarketSummary{
		Market: trade.Market,
	}
}

func (ms *MarketSummary) AddTrade(trade Trade) {
	ms.MeanPrice = (trade.Price + float64(ms.Count)*ms.MeanPrice) / float64(ms.Count+1)
	ms.MeanVolume = (trade.Volume + float64(ms.Count)*ms.MeanVolume) / float64(ms.Count+1)
	ms.VolumeWeightedAveragePrice = (trade.Price*trade.Volume + ms.VolumeWeightedAveragePrice*ms.TotalVolume) / (ms.TotalVolume + trade.Volume)
	ms.TotalVolume += trade.Volume
	if trade.IsBuy {
		ms.PercentageBuy = (1 + float64(ms.Count)*ms.PercentageBuy) / float64(ms.Count+1)
	} else {
		ms.PercentageBuy = float64(ms.Count) * ms.PercentageBuy / float64(ms.Count+1)
	}
	ms.Count++
}
