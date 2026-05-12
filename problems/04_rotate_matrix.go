package problems

func RotateMatrix(m [][]int) {

	n := len(m)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {

			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}

	for i := 0; i < n; i++ {

		left := 0
		right := n - 1

		for left < right {

			m[i][left], m[i][right] =
				m[i][right], m[i][left]

			left++
			right--
		}
	}

}
