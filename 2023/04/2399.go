package leecode

func Q2399(s string, distance []int) bool {
	type index struct {
		start int
		end   int
	}

	m := make(map[rune]index)

	for i, v := range s {
		if _, ok := m[v]; ok {
			m[v] = index{m[v].start, i}
			dis := m[v].end - 1 - m[v].start
			if distance[int(v)-97] != dis {
				return false
			}
		} else {
			m[v] = index{i, -1}
		}
	}

	return true
}

func Q2399_1(s string, distance []int) bool {
	for index := range s {
		site := s[index] - 'a'
		if distance[site] == -1 {
			continue
		}
		if index+distance[site]+1 >= len(s) || s[index] != s[index+distance[site]+1] {
			return false
		}
		distance[site] = -1
	}
	return true
}
