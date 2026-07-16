func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1 

	for (i < len(numbers)) && (i < j) {
		a := numbers[i]
		b := numbers[j]
		sum := a + b 
		if sum == target {
			return []int{i+1, j+1} // 1 indexed return
		} else if sum < target {
			// we need to go bigger, move i to the right
			i = i+1
		} else {
			j = j-1
		}
	}

	// should never hit this 
	badVal := []int{-1,-1}
	return badVal
}
