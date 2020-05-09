package array_list

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	arr := NewArray(10)
	arr.Add(11)
	arr.Add(12)
	arr.Add(22)
	if !reflect.DeepEqual(arr.GetAll(), []int{11, 12, 22}) {
		t.Log(arr.GetAll())
		t.Error("error:添加失败")
	}
	arr.Add(13)
	arr.Add(14)
	arr.Add(15)
	if element, err := arr.Get(4); err != nil || element != 14 {
		t.Error("error:Get()错误")
	}
	arr.Remove(4)
	if !reflect.DeepEqual(arr.GetAll(), []int{11, 12, 22, 13, 15}) {
		t.Error("error:Remove()错误")
	}
	arr.Set(3, 100)
	if !reflect.DeepEqual(arr.GetAll(), []int{11, 12, 22, 100, 15}) {
		t.Error("error:Set()错误")
	}
}
