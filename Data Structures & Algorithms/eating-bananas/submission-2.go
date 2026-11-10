func minEatingSpeed(piles []int, h int) int {
	largest := 0
	for _, val := range piles {
		largest = max(largest, val)
	}

	left := 1
	right := largest
	res := largest

	for left <= right {
        mid := left + (right - left) / 2

        if canFinish(piles, mid, h) {
            res = mid       // This speed works, but let's see if a smaller one works
            right = mid - 1 // Try looking to the left
        } else {
            left = mid + 1  // Too slow, we need a larger speed
        }
    }

    return res
	
}

func canFinish(piles []int, k int, h int) bool{
	totalTime := 0
	for _, val := range(piles) {
		totalTime += (val + k - 1) / k // ceiling division
	}

	if totalTime <= h {
		return true
	}
	return false
}