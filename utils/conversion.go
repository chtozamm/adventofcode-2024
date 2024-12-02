package utils

import "strconv"

func StringsToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, s := range strs {
		n, _ := strconv.Atoi(s)
		ints[i] = n
	}
	return ints
}
