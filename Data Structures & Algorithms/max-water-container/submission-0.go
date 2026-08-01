
func maxArea(heights []int) int {
	highestVol := 0 
	left := 0
	right := len(heights) -1

	for left < right {
		a := heights[left]
		b := heights[right]
		vol := min(a, b) * (right - left)
		highestVol = max(highestVol, vol)

		if a < b {
			left++
		} else {right--}
	}

	return highestVol
}
