package queque

import (
	"fmt"
	"testing"
)

func TestMaxQueue(t *testing.T) {
	obj := Constructor59()
	param_1 := obj.Max_value()
	fmt.Println(param_1)
	obj.Push_back(10)
	//param_3 := obj.Pop_front()
	//fmt.Println(param_3)
	obj.Push_back(20)
	fmt.Println(obj)
}
