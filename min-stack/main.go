package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	values []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.values = append(this.values, val)
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
	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	min := math.MaxInt
	for _, v := range this.values {
		if v < min {
			min = v
		}
	}
	return min
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
