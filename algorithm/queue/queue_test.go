package queue

import "testing"

func TestQueue_enQueue(t *testing.T) {
	queue := New()
	queue.enQueue(1)
	queue.enQueue(10)
	queue.enQueue(19)
	num, err := queue.Peek()
	if err != nil {
		t.Error(err.Error())

	}
	if num != 1 {
		t.Error("queue is error")
	}
	num, err = queue.deQueue()
	if num == 1 || err != nil {
		t.Error("queue is error")
	}
}
