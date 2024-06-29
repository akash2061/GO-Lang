package custom

func Recursion(num int) int {
	if num == 0 {
		return 1
	}
	return num * Recursion(num-1)
}
