func productExceptSelfSlow(nums []int) []int {
    // slower method
	// track total product and divide by element's value 
	totalProduct := 1
	n:= len(nums)
	res := make([]int, n)
	indexesOfZeroes := []int{}
	for i := 0; i < n; i++ {
		val := nums[i]
		if val != 0 {
			totalProduct = totalProduct * val
		} else {
			indexesOfZeroes = append(indexesOfZeroes, i)
		}
	}
	if len(indexesOfZeroes) > 1 {
		//more than one zero, so ALL products-of-others is zero
		return res
	} 

	zeroIndex := -1
	if len(indexesOfZeroes) == 1 {
		zeroIndex = indexesOfZeroes[0]
		res[zeroIndex] = totalProduct
		return res
	}

	for i := 0; i < n; i++ {
		res[i] = totalProduct / nums[i]
	}

	return res
}

func productExceptSelf(nums []int) []int {
	// faster method
	// prefix/suffix tracking
	n := len(nums)
	prefix := make([]int, n)
	suffix := make([]int, n)
	result := make([]int, n)
	prefix[0] = 1
	suffix[n-1] = 1
	for i := 1; i <= n-1; i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}
	for i:= n-2; i >= 0; i-- {
		suffix[i] = nums[i+1] * suffix[i+1]
	}
	
	for i := 0; i < n; i++ {
		result[i] = prefix[i] * suffix[i]
	}

	return result
}