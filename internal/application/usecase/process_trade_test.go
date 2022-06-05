package usecase

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/jpdejavite/messari-THA-solution/internal/domain/entity"
)

func Test_Execute(t *testing.T) {

	type args struct {
		trades []entity.Trade
	}

	tests := []struct {
		name string
		args args
		want map[int]entity.MarketSummary
	}{
		{
			name: "multiple trades",
			args: args{
				trades: []entity.Trade{
					{
						ID:     6,
						Market: 8,
						Price:  7.37,
						Volume: 2633.63,
						IsBuy:  true,
					},
					{
						ID:     7,
						Market: 7,
						Price:  10,
						Volume: 100,
						IsBuy:  true,
					},
					{
						ID:     8,
						Market: 7,
						Price:  20,
						Volume: 200,
						IsBuy:  false,
					},
					{
						ID:     9,
						Market: 7,
						Price:  30,
						Volume: 200,
						IsBuy:  true,
					},
					{
						ID:     10,
						Market: 7,
						Price:  40,
						Volume: 500,
						IsBuy:  false,
					},
				},
			},
			want: map[int]entity.MarketSummary{
				7: {
					Market:                     7,
					TotalVolume:                1000,
					MeanPrice:                  25,
					MeanVolume:                 250,
					VolumeWeightedAveragePrice: 31,
					PercentageBuy:              0.5,
					Count:                      4,
				},
				8: {
					Market:                     8,
					TotalVolume:                2633.63,
					MeanPrice:                  7.37,
					MeanVolume:                 2633.63,
					VolumeWeightedAveragePrice: 7.37,
					PercentageBuy:              1,
					Count:                      1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := NewProcessTrade()

			for _, trade := range tt.args.trades {
				b, _ := json.Marshal(trade)
				pt.Execute(string(b))
			}

			if !reflect.DeepEqual(pt.GetTradesSummary(), tt.want) {
				t.Errorf("Execute() = %v, want %v", pt.GetTradesSummary(), tt.want)
			}
		})
	}
}
