package problems

func RotateStringWithGoal(s string, goal string) bool {
	n := len(s)
	if n == 0 {
		return false
	}

	for k := range n {
		rot := s[n-k:] + s[:n-k]
		if rot == goal {
			return true
		}
	}
	return false

}
