package main

import "fmt"

func main() {
	// nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	nums := []int{1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	fmt.Println(longestOnes(nums))
}

func longestOnes(nums []int) int {
	k := 1
	begin := 0
	windowState := 0
	result := 0

	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			windowState++
		}

		for windowState > k {
			if nums[begin] == 0 {
				windowState -= 1
			}

			begin += 1 // shrink window
		}

		result = max(result, end-begin+1)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
