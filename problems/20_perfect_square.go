package problems

func IsPerfectSquare(n int) bool {
	if n < 1 {
		return false
	}

	l, r := 1, n

	for l <= r {
		mid := (l + r) / 2
		sq := mid * mid

		if sq == n {
			return true
		} else if sq < n {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
