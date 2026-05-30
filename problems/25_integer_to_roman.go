package problems

// 3749 -> MMMDCCXLIX -> MMM = 3000, D = 500, CC= 100+100 = 200, XL = 40, IX= -1 + 10 = 9 = 3000 + 500 + 200 + 40 + 9 = 3749

func IntegerToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	n := len(values) // 0,1,2,3 = 4
	result := ""     // first iteration ->  MMM, second iteration -> D, third interation -> CC, 7 () iteration -> XL, final iteration -> IX
	// MMMDCCXLIX

	for i := range n {
		for num >= values[i] { // num = 9 >= values[9] = 9
			result += romans[i] // result = romans[9] ->  IX
			num -= values[i]    // values[9] -> 9, = num = 9 -  9 = 0
		}
	}

	return result
}
