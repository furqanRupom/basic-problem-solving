package problems

func CheckPalindrome(list []int) bool {
	first := 0
	last := len(list) - 1
	for first < last {
		if list[first] != list[last] {
			return false
		}
		first++
		last--
	}
	return true
}
