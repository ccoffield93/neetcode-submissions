func productExceptSelf(nums []int) []int {
	prefixArray := make([]int, len(nums))
	suffixArray := make([]int, len(nums))
	output := make([]int, len(nums))

    prefixArray[0] = 1
	suffixArray[len(nums) - 1] = 1
	for i:= 1; i < len(nums); i++ {
		prefixArray[i] = prefixArray[i-1]*nums[i-1]
	}

	for i:= len(nums) - 2; i >= 0; i-- {
		suffixArray[i] = suffixArray[i+1]*nums[i+1]
	}

	for i:= 0; i < len(nums); i++ {
		output[i] = prefixArray[i] * suffixArray[i]
	}

	return output
}
