package entity

import (
	"reflect"
	"testing"
)

func Test_AddTrade(t *testing.T) {

	type args struct {
		trades []Trade
	}

	tests := []struct {
		name string
		args args
		want MarketSummary
	}{
		{
			name: "single trade",
			args: args{
				trades: []Trade{
					{
						ID:     7,
						Market: 7,
						Price:  7.37,
						Volume: 2633.63,
						IsBuy:  true,
					},
				},
			},
			want: MarketSummary{
				Market:                     7,
				TotalVolume:                2633.63,
				MeanPrice:                  7.37,
				MeanVolume:                 2633.63,
				VolumeWeightedAveragePrice: 7.37,
				PercentageBuy:              1,
				Count:                      1,
			},
		},
		{
			name: "multiple trade",
			args: args{
				trades: []Trade{
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
			want: MarketSummary{
				Market:                     7,
				TotalVolume:                1000,
				MeanPrice:                  25,
				MeanVolume:                 250,
				VolumeWeightedAveragePrice: 31,
				PercentageBuy:              0.5,
				Count:                      4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMarketSummary(tt.args.trades[0])

			for _, trade := range tt.args.trades {
				got.AddTrade(trade)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTrade() = %v, want %v", got, tt.want)
			}
		})
	}
}
