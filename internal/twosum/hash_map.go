package twosum

func HashMap(nums []int, target int) []int {
	seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {

		t := target - nums[i]
		_, exists := seen[t]

		if exists {
			return []int{seen[t], i}
		} else {
			seen[nums[i]] = i
		}

	}

	return nil
}
