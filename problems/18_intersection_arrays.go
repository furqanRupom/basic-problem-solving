package problems

func Intersection(a, b []int) []int {
	set := map[int]bool{}
	res := []int{}

	for _, v := range a {
		set[v] = true
	}

	for _, v := range b {
		if set[v] {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}
