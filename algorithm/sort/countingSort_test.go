package sort

import (
	"reflect"
	"testing"
)

func Test_countingSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "计数排序",
			args: args{
				nums: []int{1, 4, 5, 6, -99, 99, 99},
			},
			want: []int{-99, 1, 4, 5, 6, 99, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countingSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countingSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
