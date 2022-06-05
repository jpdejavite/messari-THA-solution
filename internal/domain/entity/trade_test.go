package entity

import (
	"reflect"
	"testing"
)

func Test_NewTrade(t *testing.T) {

	type args struct {
		text string
	}

	tests := []struct {
		name string
		args args
		want *Trade
	}{
		{
			name: "valid json",
			args: args{
				text: `{"id":7,"market":7,"price":7.37,"volume":2633.63,"is_buy":false}`,
			},
			want: &Trade{
				ID:     7,
				Market: 7,
				Price:  7.37,
				Volume: 2633.63,
				IsBuy:  false,
			},
		},
		{
			name: "invalid json",
			args: args{
				text: "INVALID",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTrade(tt.args.text)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTrade() = %v, want %v", got, tt.want)
			}
		})
	}
}
