package queue

import "errors"

//用slice实现队列

type Queue struct {
	queue []int
	len   int
}

func New() *Queue {
	return &Queue{
		queue: make([]int, 0),
		len:   0,
	}
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) enQueue(element int) {
	q.queue = append(q.queue, element)
	q.len++
	return
}

func (q *Queue) deQueue() (int, error) {
	if q.len <= 0 {
		return 0, errors.New("stack is empty")
	}
	num := q.queue[0]
	q.queue = q.queue[1:]
	q.len--
	return num, nil
}

func (q *Queue) Peek() (int, error) {
	if q.len <= 0 {
		return 0, errors.New("stack is empty")
	}
	return q.queue[0], nil
}
