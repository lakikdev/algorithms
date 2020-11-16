package sol1

func TwoNums(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for y := i+1; y < len(nums); y++ {
			if nums[i] + nums[y] == target {
				return []int{i, y}
			}
		}
	}
	return []int{-1, -1}
}