package sol3

func TwoNums(nums []int, target int) []int {

	dict := make(map[int]int, 0)

	for index, num := range nums {
		if index2, found := dict[target-num]; found {
			return []int{index2, index}
		}
		dict[num] = index
	}

	return []int{-1, -1}
}