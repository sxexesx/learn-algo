package main

import "fmt"

func main() {
	// prices := []int{6, 7, 1, 5, 3, 6, 4}
	prices := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println("max profit: ", maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min := prices[0]
	var result int

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}

		if profit := prices[i] - min; profit > result {
			result = profit
		}
	}

	return result
}
