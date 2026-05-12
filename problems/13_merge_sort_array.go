package problems

func MegetSortedArray(a, b []int) []int {
	i, j := 0, 0
	res := []int{}

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	res = append(res, a[i:]...)
	res = append(res, b[j:]...)

	return res
}
