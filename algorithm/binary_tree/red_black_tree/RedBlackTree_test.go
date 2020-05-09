package red_black_tree

import (
	"reflect"
	"testing"
)

func testTree(nums []int64) *Tree {
	tree := NewTree()
	for _, v := range nums {
		tree.Add(v)
	}

	return tree
}
func TestTree_LevelOrderTraversal(t1 *testing.T) {
	nums1 := []int64{81, 34, 43, 78, 7, 62, 71, 40, 48, 5, 83, 90, 59, 92, 8, 84, 51}
	tests := []struct {
		name   string
		fields *Tree
		want   []NodeValue
	}{
		// TODO: Add test cases.
		{
			name:   "中序遍历",
			fields: testTree(nums1),
			want: []NodeValue{
				{43, true},
				{34, true},
				{78, true},
				{7, true},
				{40, true},
				{62, false},
				{83, false},
				{5, false},
				{8, false},
				{51, true},
				{71, true},
				{81, true},
				{90, true},
				{48, false},
				{59, false},
				{84, false},
				{92, false}},
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
	nums := []int64{61, 97, 14, 49, 29, 9}
	type fields struct {
		root *Node
		size int
	}
	type args struct {
		element int64
	}
	tests := []struct {
		name     string
		fields   *Tree
		args     args
		want     bool
		wantList []NodeValue
	}{
		// TODO: Add test cases.
		{
			name:   "层序遍历2",
			args:   args{element: 49},
			fields: testTree(nums),
			wantList: []NodeValue{
				{61, BLACK},
				{14, RED},
				{97, BLACK},
				{9, BLACK},
				{29, BLACK}},
			want: true,
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
