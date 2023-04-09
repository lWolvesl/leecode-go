package leecode

import "strconv"

func Q1017(n int) string {
	if n == 0 {
		return "0"
	}

	binary := ""
	for n != 0 {
		remainder := n % -2
		n = n / -2
		if remainder < 0 {
			remainder += 2
			n++
		}
		binary = strconv.Itoa(remainder) + binary
	}

	return binary
}
