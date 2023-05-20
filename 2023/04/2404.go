package leecode

func mostFrequentEven(nums []int) int {
	ret := -1
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums) && nums[i]%2 == 0; i++ {
		c, ok := m[nums[i]]
		if !ok {
			m[nums[i]] = 1
		} else {
			m[nums[i]] = c + 1
		}
		if ret == -1 {
			ret = i
			continue
		}
		if m[nums[i]] == m[nums[ret]] {
			if nums[i] < nums[ret] {
				ret = i
			}
		} else if m[nums[i]] > m[nums[ret]] {
			ret = i
		}
	}
	if ret == -1 {
		return -1
	}
	return nums[ret]
}
