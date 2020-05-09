package _map

import "testing"

func Test_hashCode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "计算hashcode",
			args: args{str: "jack"},
			want: 3254239,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashCode(tt.args.str); got != tt.want {
				t.Errorf("hashCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
