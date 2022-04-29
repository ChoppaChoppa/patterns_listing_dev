package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "RemoveDup_test_1",
			args: args{
				arr: []string{"a", "b", "c", "b", "a"},
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicate(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCase(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "LoserCase_test_1",
			args: args{
				[]string{"a", "b", "A", "c", "C"},
			},
			want: []string{"a", "b", "a", "c", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LowerCase(tt.args.arr)
			if got := tt.args.arr; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveSpacesAtTheEnd(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "RemoveSpaces_test_1",
			args: args{
				arr: []string{"asdfe", "rigjoej", "asodijo   ", "asdasd ", "asodiaosi j       "},
			},
			want: []string{"asdfe", "rigjoej", "asodijo", "asdasd", "asodiaosi j"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveSpacesAtTheEnd(tt.args.arr)

			if got := tt.args.arr; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
