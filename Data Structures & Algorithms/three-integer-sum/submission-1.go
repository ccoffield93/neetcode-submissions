import "slices"
func threeSum(nums []int) [][]int {
	// start with a sorted array so we can use 2-pointer 
	slices.Sort(nums)

	toReturn := [][]int{}
	i := 0
	j := 1
	k := len(nums) - 1

	// we're basically doing 2-sum with each possible initial value
	// and all elements after it
	for i = 0; i <= len(nums) - 3; i++ {
		// if the value at i is the same as its predecessor, 
		// we'd get duplicate results, so skip it
		if i > 0 && nums[i] == nums [i-1] {
			continue
		}

		// start looking at everything AFTER i 
		j = i + 1
		k = len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				// save this one, move j right after to keep looking for combinations
				toReturn = append(toReturn, []int{nums[i],nums[j],nums[k]})
				j++ 
				// if the next j's value is identical to the last, we'd get a duplicate
				// move j to the right again 
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if sum < 0 {
				// we need bigger numbers, move j right
				j++
			} else if sum > 0 {
				// we need smaller numbers, move k left
				k--
			}
		}
	}

	return toReturn
}
