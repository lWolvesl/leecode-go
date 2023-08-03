package lee2308

import "sort"

func sumOfPower(nums []int) int {
	sort.IntSlice(nums).Sort()
	sum := 0
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		for j := 0; j < i; j++ {
			dis := i - j
			co := (dis*dis + dis) / 2 * nums[j]
			temp += co
		}
		sum += temp * nums[i]
	}
	return sum
}
