package main

import "testing"

func test2427(t *testing.T) {
	t.Log(commonFactors(24, 12))
}

func commonFactors(a int, b int) int {
	count := 0
	arr := make([]int, 3)
	if a > b {
		arr[0] = a
		arr[1] = b
	} else {
		arr[0] = a
		arr[1] = b
	}
	for i := 1; i < count; i++ {

	}
	return count
}
