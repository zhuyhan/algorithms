package stack

import "testing"

func TestStack_Top(t *testing.T) {
	type fields struct {
		element []int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "获取栈顶元素",
			fields:  fields{element: []int{1, 2, 4, 99}},
			want:    99,
			wantErr: false,
		},
		{
			name:    "获取栈顶元素-空",
			fields:  fields{element: []int{}},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				element: tt.fields.element,
			}
			got, err := s.Top()
			if (err != nil) != tt.wantErr {
				t.Errorf("Top() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Top() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	stack := New()
	stack.Push(1)
	stack.Push(9)
	stack.Push(19)
	stack.Push(29)
	num, err := stack.Pop()
	if err != nil {
		t.Error(err.Error())
	}
	if num != 29 {
		t.Error("err")
	}
}
