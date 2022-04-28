package main

import (
	"reflect"
	"testing"
)

//TODO: доделать тест
func TestOrChanel(t *testing.T) {
	type args struct {
		ch []<-chan interface{}
	}
	tests := []struct {
		name string
		args args
		want <-chan interface{}
	}{
		{
			name: "OrChanel_test_1",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrChanel(tt.args.ch...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrChanel() = %v, want %v", got, tt.want)
			}
		})
	}
}
