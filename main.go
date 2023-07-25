package main

import (
	"fmt"
	"leecode/dataStruct/Queue"
)

func main() {
	x := []int{4, 3, 7, 1, 8, 6}
	queue := Queue.NewPriorityQueue([]int{})
	for _, v := range x {
		queue.Push(v)
	}
	a := queue.Pop()
	fmt.Println(a)
}
