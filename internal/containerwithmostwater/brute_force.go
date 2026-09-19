package containerwithmostwater

func BruteForce(height []int) int {
	maxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			width := j - i
			h := min(height[i], height[j])
			area := width * h
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
