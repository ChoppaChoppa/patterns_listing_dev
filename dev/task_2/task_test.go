package main

import "testing"

func TestUnPackStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_1",
			args: args{
				"a4bc2d5e",
			},
			want: "aaaabccddddde",
		},
		{
			name: "test_2",
			args: args{
				"abcd",
			},
			want: "abcd",
		},
		{
			name: "test_3",
			args: args{
				"45",
			},
			want: "error",
		},
		{
			name: "test_4",
			args: args{
				"",
			},
			want: "",
		},
		{
			name: "test_5",
			args: args{
				"qwe\\4\\5",
			},
			want: "qwe45",
		},
		{
			name: "test_6",
			args: args{
				"qwe\\45",
			},
			want: "qwe44444",
		},
		{
			name: "test_7",
			args: args{
				"qwe\\\\5",
			},
			want: "qwe\\\\\\\\\\",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnPackStr(tt.args.str); got != tt.want {
				t.Errorf("UnPackageStr() = %v, want %v", got, tt.want)
			}
		})
	}
}