package binary_search_tree

import (
	"reflect"
	"testing"
)

func testTree() *Tree {
	tree1 := NewTree()
	nums := []int{53, 36, 21, 96, 51, 33, 75, 72, 66, 11, 14}
	for _, v := range nums {
		tree1.Add(v)
	}

	return tree1
}

func testTree2() *Tree {
	tree1 := NewTree()
	nums := []int{48, 16, 97, 21, 85, 15}
	for _, v := range nums {
		tree1.Add(v)
	}

	return tree1
}

//生成一个二叉树，和抛出一个结点66
func testTree3() (*Tree, *Node) {
	tree1 := NewTree()
	nums := []int{53, 36, 21, 96, 51, 33, 75, 72, 66, 11, 14}
	for _, v := range nums {
		tree1.Add(v)
	}
	return tree1, tree1.root.right.left.left.left
}

//生成一个二叉树，和抛出一个结点21
func testTree4() (*Tree, *Node) {
	tree1 := NewTree()
	nums := []int{53, 36, 21, 96, 51, 33, 75, 72, 66, 11, 14}
	for _, v := range nums {
		tree1.Add(v)
	}
	return tree1, tree1.root.left.left
}

func TestTree_PreOrderTraversal(t1 *testing.T) {
	tests := []struct {
		name   string
		fields *Tree
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "前序遍历",
			fields: testTree(),
			want:   false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := t.PreOrderTraversal(); got != tt.want {
				t1.Errorf("PreOrderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_InOrderTraversal(t1 *testing.T) {
	tests := []struct {
		name   string
		fields *Tree
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "前序遍历",
			fields: testTree(),
			want:   false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := t.InOrderTraversal(); got != tt.want {
				t1.Errorf("InOrderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_PostOrderTraversal(t1 *testing.T) {
	tests := []struct {
		name   string
		fields *Tree
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "前序遍历",
			fields: testTree(),
			want:   false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := t.PostOrderTraversal(); got != tt.want {
				t1.Errorf("PostOrderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_LevelOrderTraversal(t1 *testing.T) {
	tests := []struct {
		name   string
		fields *Tree
		want   []int
	}{
		// TODO: Add test cases.
		{
			name:   "层次遍历",
			fields: testTree(),
			want:   []int{53, 36, 96, 21, 51, 75, 11, 33, 72, 14, 66},
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

func TestTree_Height(t1 *testing.T) {
	tests := []struct {
		name   string
		fields *Tree
		want   int
	}{
		// TODO: Add test cases.
		{
			name:   "层次遍历-树高度",
			fields: testTree(),
			want:   5,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := t.Height(); got != tt.want {
				t1.Errorf("Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Height2(t1 *testing.T) {
	type fields struct {
		root *Node
		size int
	}
	type args struct {
		root *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "递归遍历行数",
			args: args{root: testTree().root},
			want: 5,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := t.Height2(tt.args.root); got != tt.want {
				t1.Errorf("Height2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_IsComplete(t1 *testing.T) {
	tests := []struct {
		name   string
		fields *Tree
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "不是完整二叉树",
			fields: testTree(),
			want:   false,
		},
		{
			name:   "完整二叉树",
			fields: testTree2(),
			want:   true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := t.IsComplete(); got != tt.want {
				t1.Errorf("IsComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Predecessor(t1 *testing.T) {
	tree1, node1 := testTree3()
	tree2, node2 := testTree4()
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields *Tree
		args   args
		want   *Node
	}{
		// TODO: Add test cases.
		{
			name:   "前驱结点",
			fields: tree1,
			args:   args{node: node1},
			want:   node1.parent.parent.parent.parent,
		},
		{
			name:   "前驱结点",
			fields: tree2,
			args:   args{node: node2},
			want:   node2.left.right,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := t.Predecessor(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Predecessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Successor(t1 *testing.T) {
	tree1, node1 := testTree3()
	tree2, node2 := testTree4()
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields *Tree
		args   args
		want   *Node
	}{
		// TODO: Add test cases.
		{
			name:   "后继",
			fields: tree1,
			args:   args{node: node1},
			want:   node1.parent,
		},
		{
			name:   "后继结点",
			fields: tree2,
			args:   args{node: node2},
			want:   node2.right,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := t.Successor(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Successor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_GetNode(t1 *testing.T) {
	type args struct {
		element int
	}
	tests := []struct {
		name   string
		fields *Tree
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name:   "查找结点",
			fields: testTree(),
			args:   args{element: 66},
			want:   66,
		},
		{
			name:   "查找结点",
			fields: testTree2(),
			args:   args{element: 48},
			want:   48,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := t.GetNode(tt.args.element); got == nil || got.element != tt.want {
				t1.Errorf("GetNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Remove(t1 *testing.T) {
	type args struct {
		element int
	}
	tests := []struct {
		name   string
		fields *Tree
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "删除",
			fields: testTree2(),
			args:   args{element: 48},
			want:   true,
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
		})
	}
}
