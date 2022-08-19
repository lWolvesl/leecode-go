package mouth8

import (
	"fmt"
)

/**
给你一个由若干 0 和 1 组成的字符串 s ，请你计算并返回将该字符串分割成两个 非空 子字符串（即 左 子字符串和 右 子字符串）所能获得的最大得分。

「分割字符串的得分」为 左 子字符串中 0 的数量加上 右 子字符串中 1 的数量。
*/
func day14(s string) int {
	max := 0
	length := len(s)
	index := 0
	for index < length-1 {
		pear := 0
		rear := 0
		for i := 0; i <= index; i++ {
			if s[i] == '0' {
				pear++
			}
		}
		for i := index + 1; i < length; i++ {
			if s[i] == '1' {
				rear++
			}
		}
		if max < pear+rear {
			max = pear + rear
		}
		index++
	}
	return max
}

//动态规划
func day14_1(s string) int {
	n := len(s)
	presum, ans := 0, -1-n
	for i := 0; i < n; i++ {
		if cur := presum*2 - i; i > 0 && cur > ans {
			ans = cur
		}
		if s[i] == '0' {
			presum++
		}
	}
	return ans + n - presum
}

func runday14() {
	s := "011101"
	fmt.Print(day14(s))
}

func day19(startTime []int, endTime []int, queryTime int) int {
	count := 0
	for i := range startTime {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			count++
		}
	}
	return count
}

func Run8() {
	fmt.Println("start")
	runday14()
}
