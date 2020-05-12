package stack

import "errors"

//用slice实现栈

type Stack struct {
	element []int
}

func New() *Stack {
	return &Stack{element: make([]int, 0)}
}

func (s *Stack) Len() int {
	return len(s.element)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

//Pop 出栈，若为空弹出-1
func (s *Stack) Pop() (int, error) {
	if s.Len() == 0 {
		return -1, errors.New("out of index")
	}
	num := s.element[s.Len()-1]
	s.element = s.element[:s.Len()-1]
	return num, nil
}

//Push 入栈
func (s *Stack) Push(element int) {
	s.element = append(s.element, element)
	return
}

//Top 栈顶元素的值
func (s *Stack) Top() (int, error) {
	if s.Len() == 0 {
		return -1, errors.New("out of index")
	}
	num := s.element[s.Len()-1]
	return num, nil
}
