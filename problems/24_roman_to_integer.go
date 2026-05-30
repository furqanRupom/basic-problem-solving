package problems

// I -> 1, III -> 3, XII -> 12, IX = 9
func RomanToInteger(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	target := 0
	n := len(s) // 0,1
	for i := range n {
		currentRoman := romanMap[s[i]] // romanMap[s[i]] -> romanMap[s[X]] -> 10

		if i+1 < len(s) && currentRoman < romanMap[s[i+1]] { // 2 < 2 && 2 < romanMap[s[]]
			target -= currentRoman // 0 - 1 = -1
		} else {
			target += currentRoman // -1 + 10 = 9
		}
	}

	return target
}
