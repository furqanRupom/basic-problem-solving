package problems

func CountVowels(s string) int {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	count := 0

	for _, ch := range s {
		if vowels[ch] {
			count++
		}
	}
	return count
}
