package sort

import (
	"reflect"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
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
			name: "冒泡排序",
			args: args{
				nums: []int{1, 4, 5, 6, 2, 3},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubbleSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectionSort(t *testing.T) {
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
			name: "选择排序",
			args: args{
				nums: []int{1, 4, 5, 6, 2, 3},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectionSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{
			name: "插入排序",
			args: args{
				nums: []int{1, 4, 5, 6, 2, 3},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shellSort(t *testing.T) {
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
			name: "希尔排序",
			args: args{
				nums: []int{1, 4, 5, 6, 2, 3, 11, 45, 30, 31, 12, 11, 23, 12},
			},
			want: []int{1, 2, 3, 4, 5, 6, 11, 11, 12, 12, 23, 30, 31, 45},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shellSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shellSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testMergeSort(t *testing.T) {
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
			name: "归并排序",
			args: args{
				nums: []int{1, 4, 5, 6, 2, 3, 11, 45, 30, 31, 12, 11, 23, 12},
			},
			want: []int{1, 2, 3, 4, 5, 6, 11, 11, 12, 12, 23, 30, 31, 45},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testMergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testMergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testQuickSort(t *testing.T) {
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
			name: "归并排序",
			args: args{
				nums: []int{1, 4, 5, 6, 2, 3, 11, 45, 30, 31, 12, 11, 23, 12},
			},
			want: []int{1, 2, 3, 4, 5, 6, 11, 11, 12, 12, 23, 30, 31, 45},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testQuickSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testQuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
