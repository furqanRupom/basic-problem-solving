package problems

func CountTotalSpecialChar(word string) int {
	ls := make(map[rune]bool)
	us := make(map[rune]bool)
	count := 0

	for _, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			ls[ch] = true
		} else if ch >= 'A' && ch <= 'Z' {
			us[ch] = true
		}
	}

	for ch := 'a'; ch <= 'z'; ch++ {
		if ls[ch] && us[ch-32] {
			count++
		}
	}

	return count
}
