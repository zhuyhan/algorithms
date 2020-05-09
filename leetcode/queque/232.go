package queque

// 232. 用栈实现队列
// https://leetcode-cn.com/problems/implement-queue-using-stacks/
type MyQueue struct {
	inStack  []int
	outStack []int
}

/** Initialize your data structure here. */
func Constructor232() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.Peek()
	e := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return e
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		//出队的栈为空，将入队的栈写入出栈的队
		for len(this.inStack) > 0 {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
	}
	return this.outStack[len(this.outStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.outStack)+len(this.inStack) == 0
}
