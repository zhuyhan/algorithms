package heap

import (
	"reflect"
	"testing"
)

func buildHeap(list []int) *MaxHeap {

	heap := Init()
	for _, v := range list {
		heap.Add(v)
	}
	return heap
}

func TestMaxHeap(t *testing.T) {
	list := []int{68, 72, 43, 50, 38, 1}
	heap := buildHeap(list)
	if !reflect.DeepEqual(heap.element, []int{72, 68, 43, 50, 38, 1}) {
		t.Error(heap.element)
	}
}

func TestMaxHeap_Remove(t *testing.T) {
	list := []int{68, 72, 43, 50, 38, 1}
	tests := []struct {
		name   string
		fields *MaxHeap
		want   []int
	}{
		// TODO: Add test cases.
		{
			name:   "删除",
			fields: buildHeap(list),
			want:   []int{68, 50, 43, 1, 38},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &MaxHeap{
				element: tt.fields.element,
			}
			if b.Remove(); !reflect.DeepEqual(b.element, tt.want) {
				t.Errorf("b.element = %v, want %v", b.element, tt.want)
			}

		})
	}
}

func TestHeapify2(t *testing.T) {
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
			args: args{list: []int{1, 99, 98, 32, 10, 100}},
			want: []int{100, 99, 98, 32, 10, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Heapify2(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Heapify2() = %v, want %v", got, tt.want)
			}
		})
	}
}
