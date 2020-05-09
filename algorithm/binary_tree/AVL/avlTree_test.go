package AVL

import (
	"reflect"
	"testing"
)

func testTree(nums []int64) *Tree {
	tree1 := NewTree()
	//nums := []int64{28, 37, 45, 39, 10, 20, 53, 42, 56, 83, 38, 33}
	for _, v := range nums {
		tree1.Add(v)
	}

	return tree1
}

func TestTree_LevelOrderTraversal(t1 *testing.T) {
	nums1 := []int64{28, 37, 45, 39, 10, 20, 53, 42, 56, 83, 38, 33}
	nums2 := []int64{60, 34, 53, 62, 48, 58, 65, 2, 29, 9, 32, 88, 15, 72, 30}

	tests := []struct {
		name   string
		fields *Tree
		want   []int64
	}{
		// TODO: Add test cases.
		{
			name:   "层序遍历1",
			fields: testTree(nums1),
			want:   []int64{37, 20, 45, 10, 28, 39, 56, 33, 38, 42, 53, 83},
		},
		{
			name:   "层序遍历2",
			fields: testTree(nums2),
			want:   []int64{53, 29, 65, 9, 34, 60, 88, 2, 15, 32, 48, 58, 62, 72, 30},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := t.LevelOrderTraversal(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("LevelOrderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Remove(t1 *testing.T) {
	nums2 := []int64{36, 28, 100, 25, 79, 110, 101, 111, 112}
	type args struct {
		element int64
	}
	tests := []struct {
		name     string
		fields   *Tree
		args     args
		wantList []int64
		want     bool
	}{
		// TODO: Add test cases.

		{
			name:     "层序遍历2",
			args:     args{element: 28},
			fields:   testTree(nums2),
			wantList: []int64{110, 36, 111, 25, 100, 112, 79, 101},
			want:     true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := t.Remove(tt.args.element); got != tt.want {
				t1.Errorf("Remove() = %v, want %v", got, tt.want)
			}

			if got := t.LevelOrderTraversal(); !reflect.DeepEqual(got, tt.wantList) {
				t1.Errorf("LevelOrderTraversal() = %v, want %v", got, tt.wantList)
			}

		})
	}
}
