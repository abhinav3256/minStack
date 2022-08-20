package main

import "fmt"

type MinStack struct {
	stackTop  int
	stackData []int

	minstackTop  int
	minstackData []int
}

func main() {
	result := Constructor()

	fmt.Println(result)
}
func Constructor() MinStack {
	stack := MinStack{}
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	min := stack.GetMin()
	fmt.Println(min)
	stack.Pop()

	top := stack.Top()
	fmt.Println(top)
	min = stack.GetMin()
	fmt.Println(min)

	return stack
}

func (this *MinStack) Push(val int) {
	this.stackData = append(this.stackData, val)
	this.stackTop++

	if this.minstackTop > 0 {
		if this.GetMin() > val {
			this.minstackData = append(this.minstackData, val)
			this.minstackTop++
		}
	} else {

		this.minstackData = append(this.minstackData, val)
		this.minstackTop++

	}

}

func (this *MinStack) Pop() {

	if this.minstackTop == this.stackData[this.stackTop-1] {
		this.minstackData = append(this.minstackData[:this.minstackTop-1], this.minstackData[this.minstackTop-1+1:]...)
		this.minstackTop--
	}
	this.stackData = append(this.stackData[:this.stackTop-1], this.stackData[this.stackTop-1+1:]...)
	this.stackTop--

}

func (this *MinStack) Top() int {
	return this.stackData[this.stackTop-1]
}

func (this *MinStack) GetMin() int {
	if this.minstackTop != 0 {
		return this.minstackData[this.minstackTop-1]
	}
	return -1

}
