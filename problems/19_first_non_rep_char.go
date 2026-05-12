package problems

func FirstUnique(s string) rune {
	count := map[rune]int{}

	for _, ch := range s {
		count[ch]++
	}

	for _, ch := range s {
		if count[ch] == 1 {
			return ch
		}
	}
	return -1
}
