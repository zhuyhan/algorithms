package queque

import "fmt"

//面试题59 - II. 队列的最大值
//https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/

/**
解题思路：
Step 1. 用切片模拟队列入队和出队都能做到时间复杂度O(1)
Step 2. 最大值用一个辅助切片实现；每次push新元素，先判断前面的元素有没小于该元素；
        如果有删除小于的元素，添加新元素
*/
type MaxQueue struct {
	queue []int
	max   []int
}

func Constructor59() MaxQueue {
	return MaxQueue{
		queue: make([]int, 0),
		max:   make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	//入队
	this.queue = append(this.queue, value)
	//维护最大值队列
	if len(this.max) != 0 {
		fmt.Println("this.max", this.max)
		for i := len(this.max) - 1; i >= 0; i-- {
			if this.max[i] < value {
				this.max = this.max[:i]
			} else {
				break
			}
		}
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	//出队
	if len(this.max) == 0 {
		return -1
	}
	value := this.queue[0]
	this.queue = this.queue[1:]
	//维护最大值队列
	if len(this.max) != 0 {
		if this.max[0] == value {
			this.max = this.max[1:]
		}
	}
	return value
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
