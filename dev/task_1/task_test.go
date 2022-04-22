package main

import (
	"github.com/beevik/ntp"
	"testing"
)

func TestGetTime(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "wrong_host",
			args: args{
				"00.beevik-ntp.pool.ntp.org",
			},
		},
		{
			name: "true_host",
			args: args{
				"0.beevik-ntp.pool.ntp.org",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ntp.Time(tt.args.host)
			if err != nil {
				t.Errorf("err: %v", err)
			}
		})
	}
}
