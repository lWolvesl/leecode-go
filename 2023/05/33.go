package lee2305

import "math"

func storeWater(bucket []int, vat []int) int {
	n := len(bucket)
	maxk := 0
	for _, v := range vat {
		if v > maxk {
			maxk = v
		}
	}
	if maxk == 0 {
		return 0
	}
	res := math.MaxInt32
	for k := 1; k <= maxk && k < res; k++ {
		t := 0
		for i := 0; i < n; i++ {
			t += max(0, (vat[i]+k-1)/k-bucket[i])
		}
		res = min(res, t+k)
	}
	return res
}
