func maxProfit(prices []int) int {
   l := 0
   r := 1
   maxProfit := 0
   
   for r < len(prices) {
		buy := prices[l]
		sell := prices[r]

		maxProfit = max(maxProfit, sell - buy)

		if sell < buy {
			l = r
		}

		r++
   }

   return maxProfit
}
