package problems

func SumInArray(numList []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range numList {
		needed := target - num

		if idx, ok := seen[needed]; ok {
			return []int{numList[idx], numList[i]}
		}
		seen[num] = i
	}
	return nil
}
