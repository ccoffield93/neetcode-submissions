func findMin(nums []int) int {
	// edge case: it isn't rotated at all 
	n := len(nums)
	if nums[n-1] > nums[0] || n == 1 {
		// it isn't rotated at all 
		return nums[0]
	}


	left := 1
	right := n-1 
	for left <= right {
		mid := (left + right) / 2 

		valAtMid := nums[mid]
		prevVal := nums[mid-1]

		if prevVal > valAtMid {
			// found the rotation point
			return valAtMid
		}

		// otherwise we need to know if we examine the left half
		// or the right half 
		// check 'r' to see if it's in the sorted segment 
		valAtR := nums[right]
		if valAtMid > valAtR {
			// the flip point is in the right half 
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
