package singleCircleLinked

import (
	"reflect"
	"testing"
)

func testList(nums []int) *List {
	list := New()
	for _, v := range nums {
		list.Add(v)
	}
	return list
}

func TestList_Node(t *testing.T) {
	list := testList([]int{1, 2, 4, 6, 8})
	list2 := testList([]int{1, 2, 4, 6, 8})
	list2.Insert(3, 88)
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  *List
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "list索引0",
			fields:  list,
			args:    args{index: 0},
			want:    1,
			wantErr: false,
		},
		{
			name:    "list索引1",
			fields:  list,
			args:    args{index: 1},
			want:    2,
			wantErr: false,
		},
		{
			name:    "list索引4",
			fields:  list,
			args:    args{index: 4},
			want:    8,
			wantErr: false,
		},
		{
			name:    "list2索引3",
			fields:  list2,
			args:    args{index: 3},
			want:    88,
			wantErr: false,
		},
		{
			name:    "list2索引4",
			fields:  list2,
			args:    args{index: 4},
			want:    6,
			wantErr: false,
		},
		{
			name:    "索引51",
			fields:  list,
			args:    args{index: 51},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				root: tt.fields.root,
				len:  tt.fields.len,
			}
			got, err := l.Node(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Node() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !reflect.DeepEqual(got.value, tt.want) {
				t.Errorf("Node() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_GetAll(t *testing.T) {
	list := testList([]int{1, 2, 4, 6, 8})
	list.Insert(3, 88)
	list2 := testList([]int{1, 2, 4, 6, 8})
	list2.Insert(0, 100)
	tests := []struct {
		name   string
		fields *List
		want   []int
	}{
		// TODO: Add test cases.
		{
			name:   "添加",
			fields: list,
			want:   []int{1, 2, 4, 88, 6, 8},
		},
		{
			name:   "添加",
			fields: list2,
			want:   []int{100, 1, 2, 4, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				root: tt.fields.root,
				len:  tt.fields.len,
			}
			if got := l.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Remove(t *testing.T) {
	list := testList([]int{1, 2, 4, 6, 8})
	list2 := testList([]int{19, 2, 41, 6, 8})

	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  *List
		args    args
		want    []int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "移除",
			fields:  list,
			args:    args{index: 3},
			want:    []int{1, 2, 4, 8},
			wantErr: false,
		},
		{
			name:    "移除",
			fields:  list2,
			args:    args{index: 0},
			want:    []int{2, 41, 6, 8},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				root: tt.fields.root,
				len:  tt.fields.len,
			}
			err := l.Remove(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			wantList := l.GetAll()
			if !reflect.DeepEqual(wantList, tt.want) {
				t.Errorf("GetAll() error = %v, wantErr %v", wantList, tt.want)

			}
		})
	}
}
