package heap

import (
	"reflect"
	"testing"
)

func TestTopK(t *testing.T) {
	type args struct {
		list []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "Top k问题",
			args: args{
				list: []int{1, 2, 3, 4, 6, 8, 9, 10, 99, 88, 87, 76, 78, 50, 55, 22, 44, 43, 32},
				k:    5,
			},
			want: []int{76, 78, 87, 99, 88},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopK(tt.args.list, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopK() = %v, want %v", got, tt.want)
			}
		})
	}
}
