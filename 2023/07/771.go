package lee2307

func numJewelsInStones(jewels string, stones string) int {
	j := make(map[int32]int, len(jewels))
	for _, v := range jewels {
		j[v] = 1
	}
	count := 0
	for _, v := range stones {
		if j[v] == 1 {
			count++
		}
	}
	return count
}
