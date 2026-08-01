func maxProfit(prices []int) int {
	maxProfit := 0
	left := 0
	right := 1

	if len(prices) < 2 {
		return 0 
	}

	for right < len(prices) {
		profit := prices[right] - prices[left]
		maxProfit = max(profit, maxProfit)

		if ( prices[right] < prices[left]) {
			left = right
		}
		right++
	}


	return maxProfit
}
