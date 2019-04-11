package arrays

func MaxProfitOnce(prices []int) int {
	maxProfit, minPrice := 0, 1<<32-1
	for _, price := range prices {
		if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
		if minPrice > price {
			minPrice = price
		}
	}
	return maxProfit
}
