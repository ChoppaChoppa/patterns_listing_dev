package main

import (
	"reflect"
	"regexp"
	"testing"
)

func TestBefore(t *testing.T) {
	r, _ := regexp.Compile("h")

	type args struct {
		arr []string
		reg *regexp.Regexp
		num int
	}
	tests := []struct {
		name string
		args args
		want map[int]string
	}{
		{
			name: "Before_test_1",
			args: args{
				arr: []string{"qwe", "wer", "tert", "werwer", "zxczxc", "asdfgh", "1we12e"},
				reg: r,
				num: 3,
			},
			want: map[int]string{
				5: "asdfgh",
				4: "zxczxc",
				3: "werwer",
				2: "tert",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Before(tt.args.arr, tt.args.reg, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Before() = %v, want %v", got, tt.want)
			}
		})
	}
}
