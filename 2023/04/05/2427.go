package leecode

import "log"

func Q2427(a int, b int) int {
	count := 0
	arr := make([]int, 3) //0 large 1 small 2 temp
	if a > b {
		arr[0] = a
		arr[1] = b
	} else {
		arr[0] = b
		arr[1] = a
	}
	for i := 1; i*i <= arr[1]; i++ {
		if arr[1]%i == 0 {
			arr[2] = arr[1] / i
			log.Println(arr[2])
			if arr[0]%i == 0 {
				count++
			}
			if i != arr[2] && arr[0]%arr[2] == 0 {
				count++
			}
		}
	}
	return count
}

func Q2427_1(a int, b int) int {
	g := gcd(a, b)
	ans := 0
	for i := 1; i*i <= g; i++ {
		if g%i == 0 {
			ans++
			if i*i < g {
				ans++
			}
		}
	}
	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
