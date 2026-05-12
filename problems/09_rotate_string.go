package problems

func RotateString(s string, k int) string {
	n := len(s)
	k %= n
	return s[n-k:] + s[:n-k]
}
