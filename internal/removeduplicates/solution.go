package removeduplicates

func RemoveDuplicates(nums []int) int {
	j := 0
	l := len(nums)

	for i := 0; i < l; i++ {
		if i < l-1 && nums[i] == nums[i+1] {
			continue

		}
		nums[j] = nums[i]
		j++
	}
	nums = nums[0:j]
	return j
}
