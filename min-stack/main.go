package main

import (
	"fmt"
)

type MinStack struct {
	values [][]int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) last() []int {
	return this.values[len(this.values)-1]
}

func (this *MinStack) lastVal() int {
	return this.values[len(this.values)-1][0]
}

func (this *MinStack) lastMin() int {
	return this.values[len(this.values)-1][1]
}

func (this *MinStack) Push(val int) {
	if len(this.values) == 0 {
		this.values = append(this.values, []int{val, val})
	} else {
		if val < this.lastMin() {
			this.values = append(this.values, []int{val, val})
		} else {
			this.values = append(this.values, []int{val, this.lastMin()})
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.values) == 0 {
		return
	}
	this.values = this.values[:len(this.values)-1]
}

func (this *MinStack) Top() int {
	if len(this.values) == 0 {
		return 0
	}
	return this.lastVal()
}

func (this *MinStack) GetMin() int {
	return this.lastMin()
}

func main() {
	obj := Constructor()
	obj.Push(2)
	obj.Push(-2)
	obj.Push(-1)
	obj.Push(3)
	fmt.Println(obj.GetMin())
	fmt.Println(obj.Top())
	obj.Pop()
	fmt.Println(obj.Top())
}
