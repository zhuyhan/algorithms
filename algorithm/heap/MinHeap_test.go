package heap

import (
	"reflect"
	"testing"
)

func TestHeapify(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "批量建堆",
			args: args{list: []int{1, 99, 98, 32, 10, 100, 13}},
			want: []int{1, 10, 13, 32, 99, 100, 98},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Heapify(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Heapify() = %v, want %v", got, tt.want)
			}
		})
	}
}
