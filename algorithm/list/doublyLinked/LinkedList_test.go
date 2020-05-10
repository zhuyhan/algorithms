package doublyLinked

import (
	"reflect"
	"testing"
)

func testLinked(list []interface{}) *List {
	linked := New()
	for _, v := range list {
		linked.Add(v)
	}

	return linked
}

func TestList_Node(t *testing.T) {
	nums := []interface{}{1, 2, 3, 4}
	list := testLinked(nums)
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields *List
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
		{
			name:   "index所在的位置0",
			fields: list,
			args:   args{0},
			want:   1,
		},
		{
			name:   "index所在的位置1",
			fields: list,
			args:   args{1},
			want:   2,
		},
		{
			name:   "index所在的位置7",
			fields: list,
			args:   args{7},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				first: tt.fields.first,
				last:  tt.fields.last,
				len:   tt.fields.len,
			}
			if got := l.Node(tt.args.index); got != nil && !reflect.DeepEqual(got.Value, tt.want) {
				t.Errorf("Node() = %v, want %v", got.Value, tt.want)
			}
		})
	}
}

func TestList_GetAll(t *testing.T) {
	nums := []interface{}{1, 3, 4, 5}
	list := testLinked(nums)
	list, _ = list.Insert(3, 99)
	tests := []struct {
		name   string
		fields *List
		want   []interface{}
	}{
		// TODO: Add test cases.
		{
			name:   "index所在的位置0",
			fields: testLinked(nums),
			want:   []interface{}{1, 3, 4, 5},
		},
		{
			name:   "遍历+insert",
			fields: list,
			want:   []interface{}{1, 3, 4, 99, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				first: tt.fields.first,
				last:  tt.fields.last,
				len:   tt.fields.len,
			}
			if got := l.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
