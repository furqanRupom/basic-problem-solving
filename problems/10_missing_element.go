package problems

func FindMissing(nums []int) []int {

	missing := []int{}

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		for j := 1; j < diff; j++ {
			missing = append(missing, nums[i-1]+j)
		}
	}
	return missing
}
