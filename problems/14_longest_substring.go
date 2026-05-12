package problems

func LengthOfLongestSubstring(s string) int {
	seen := map[rune]int{}
	left := 0
	maxLen := 0
	runes := []rune(s)

	for right, ch := range runes {
		if idx, ok := seen[ch]; ok && idx >= left {
			left = idx + 1
		}
		seen[ch] = right

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
