package app

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    Action
		wantErr bool
	}{
		{
			name: "Heartbeat",
			args: args{
				line: "10",
			},
			want: Action{
				Timestamp: 10,
				Type:      ACTION_HEARTBEAT,
			},
		},
		{
			name: "Sell",
			args: args{
				line: "10|1|SELL|toaster_1|10.00|20",
			},
			want: Action{
				Timestamp: 10,
				Type:      ACTION_SELL,
				Data:      []string{"10", "1", "SELL", "toaster_1", "10.00", "20"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLine(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}
