package sol2

func TwoNums(nums []int, target int) []int {

	dict := make(map[int]int, 0)

	for index, num := range nums {
		dict[num] = index
	}

	for index, num := range nums {
		if index2, found := dict[target-num]; found && index2 != index {
			return []int{index, index2}
		}
	}

	return []int{-1, -1}
}