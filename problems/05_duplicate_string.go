package problems

func RemoveDuplications(s string) string {
	seen := make(map[rune]bool)
	result := []rune{}

	for _, ch := range s {
		if !seen[ch] {
			seen[ch] = true
			result = append(result, ch)
		}
	}

	return string(result)
}
