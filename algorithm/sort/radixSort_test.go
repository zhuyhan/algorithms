package sort

import (
	"reflect"
	"testing"
)

func Test_radixSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "计数排序",
			args: args{
				nums: []int{5, 6, 1, 1199, 99, 99, 4},
			},
			want: []int{1, 4, 5, 6, 99, 99, 1199},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := radixSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("radixSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
