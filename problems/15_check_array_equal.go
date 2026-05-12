package problems

func ArraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	count := map[int]int{}

	for _, v := range a {
		count[v]++
	}
	for _, v := range b {
		count[v]--
		if count[v] < 0 {
			return false
		}
	}
	return true
}
