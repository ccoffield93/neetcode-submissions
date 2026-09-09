func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	vals := make(map[int]int)
	for i, num := range nums {
		vals[num] = i
	}

	// build a list of potential starters
	startPoints := make([]int, 0)
	for _, num := range nums {
		_, hit := vals[num-1] 
		if !hit {
			startPoints = append(startPoints, num)
		}
	}

	// using each of those start points, see how big a streak you can make
	maxStreak := 1
	for _, start := range startPoints {
		currentStreak := 1 
		currentVal := start 
		for true {
			currentVal++ 
			_, hit := vals[currentVal]
			if hit { 
				currentStreak++
			} else {
				break
			}
		}

		maxStreak = max(currentStreak, maxStreak)
	}

	return maxStreak
}
