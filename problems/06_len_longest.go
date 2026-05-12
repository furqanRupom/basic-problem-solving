package problems

func LenOfLongest(list []int) int {

	n := len(list)
	dp := make([]int, n)

	maxLen := 1

	for i := range list {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {

		for j := 0; j < i; j++ {
			if list[j] < list[i] {

				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}

		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}
