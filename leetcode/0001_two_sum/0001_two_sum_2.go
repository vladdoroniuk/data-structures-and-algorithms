package two_sum

func twoSum2(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))

	for i, num := range nums {
		complement := target - num

		if index, found := numsMap[complement]; found {
			return []int{index, i}
		}

		numsMap[num] = i
	}

	return nil
}
