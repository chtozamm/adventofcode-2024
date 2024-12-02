package utils

func AbsoluteDifference(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}
