package linked

import (
	"reflect"
	"testing"
)

func setList() *LinkList {
	list := NewLinkList()
	list.Add(10)
	list.Add(11)
	list.Add(12)
	list.Add(13)
	list.Add(14)
	return list
}

func TestLinkList_Add(t *testing.T) {

	type args struct {
		elem interface{}
	}
	tests := []struct {
		name     string
		fields   *LinkList
		args     args
		want     bool
		wantList []interface{}
	}{
		// TODO: Add test cases.
		{
			name:     "追加元素",
			fields:   NewLinkList(),
			args:     args{elem: 10},
			want:     true,
			wantList: []interface{}{10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkList{
				head: tt.fields.head,
			}
			if got := l.Add(tt.args.elem); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(l.GetAll(), tt.wantList) {
				t.Errorf("GetAll() = %v, want %v", l.GetAll(), tt.wantList)
			}
		})
	}
}

func TestLinkList_Delete(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name     string
		fields   *LinkList
		args     args
		want     bool
		wantList []interface{}
	}{
		// TODO: Add test cases.
		{
			name:     "删除",
			fields:   setList(),
			args:     args{index: 0},
			want:     true,
			wantList: []interface{}{11, 12, 13, 14},
		},
		{
			name:     "删除-索引不存在",
			fields:   setList(),
			args:     args{index: 10},
			want:     false,
			wantList: []interface{}{10, 11, 12, 13, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkList{
				head: tt.fields.head,
			}
			if got := l.Delete(tt.args.index); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(l.GetAll(), tt.wantList) {
				t.Errorf("l.GetAll() = %v, wantList %v", l.GetAll(), tt.wantList)

			}
		})
	}
}

func TestLinkList_GetData(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields *LinkList
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
		{
			name:   "获取索引1所在的值",
			fields: setList(),
			args: args{
				index: 1,
			},
			want: 11,
		},
		{
			name:   "获取索引0所在的值",
			fields: setList(),
			args: args{
				index: 0,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkList{
				head: tt.fields.head,
			}
			if got := l.GetData(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_Length(t *testing.T) {
	type fields struct {
		head *Node
	}
	tests := []struct {
		name   string
		fields *LinkList
		want   int
	}{
		// TODO: Add test cases.
		{
			name:   "表长度",
			fields: setList(),
			want:   5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkList{
				head: tt.fields.head,
			}
			if got := l.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_GetAll(t *testing.T) {

	type fields struct {
		head *Node
	}
	tests := []struct {
		name   string
		fields *LinkList
		want   []interface{}
	}{
		// TODO: Add test cases.
		{
			name:   "获取slice",
			fields: setList(),
			want:   []interface{}{10, 11, 12, 13, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkList{
				head: tt.fields.head,
			}
			if got := l.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_Insert(t *testing.T) {
	type fields struct {
		head *Node
	}
	type args struct {
		index int
		elem  interface{}
	}
	tests := []struct {
		name     string
		fields   *LinkList
		args     args
		wantList []interface{}
		want     bool
	}{
		// TODO: Add test cases.
		{
			name:   "中间插入数据",
			fields: setList(),
			args: args{
				elem:  100,
				index: 2,
			},
			want:     true,
			wantList: []interface{}{10, 11, 100, 12, 13, 14},
		},
		{
			name:   "尾部插入数据",
			fields: setList(),
			args: args{
				elem:  100,
				index: 5,
			},
			wantList: []interface{}{10, 11, 12, 13, 14, 100},
			want:     true,
		},
		{
			name:   "尾部插入数据",
			fields: setList(),
			args: args{
				elem:  100,
				index: 0,
			},
			wantList: []interface{}{100, 10, 11, 12, 13, 14},
			want:     true,
		},
		{
			name:   "索引错误",
			fields: setList(),
			args: args{
				elem:  100,
				index: 6,
			},
			wantList: []interface{}{10, 11, 12, 13, 14},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkList{
				head: tt.fields.head,
			}
			if got := l.Insert(tt.args.index, tt.args.elem); got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(l.GetAll(), tt.wantList) {
				t.Errorf("l.GetAll() = %v, wantList %v", l.GetAll(), tt.wantList)
			}
		})
	}
}
