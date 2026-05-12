package problems

func Map(arr []int, fn func(int) int) []int {
	res := make([]int, len(arr))

	for i, v := range arr {
		res[i] = fn(v)
	}
	return res
}
