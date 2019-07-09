package main

import "fmt"

type MyQueue struct {
	pushStack []int
	popStack  []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		pushStack: []int{},
		popStack:  []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.pushStack = append(this.pushStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.popStack) == 0 {
		this.TransferPushToPopStack()
	}

	popped := this.popStack[len(this.popStack)-1]
	this.popStack = this.popStack[:len(this.popStack)-1]

	return popped
}

func (this *MyQueue) TransferPushToPopStack() {
	if len(this.pushStack) == 0 {
		return
	}

	for i := len(this.pushStack) - 1; i >= 0; i-- {
		this.popStack = append(this.popStack, this.pushStack[i])
	}

	this.pushStack = this.pushStack[0:0]
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.popStack) == 0 {
		this.TransferPushToPopStack()
	}

	return this.popStack[len(this.popStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.pushStack) == 0 && len(this.popStack) == 0
}

func test() {
	obj := Constructor()
	obj.Push(5)
	obj.Push(2)
	param_2 := obj.Pop()
	param_4 := obj.Empty()

	fmt.Printf("%v, %v\n", param_2, param_4)
}

func main() {
	test()
}
