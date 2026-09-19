package containerwithmostwater

func TwoPointer(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		width := right - left

		h := height[left]
		if height[right] < h {
			h = height[right]
		}

		area := width * h

		if area > maxArea {
			maxArea = area
		}

		// Move the shorter line.
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return maxArea
}
