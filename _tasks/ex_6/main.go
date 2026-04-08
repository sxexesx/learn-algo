package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, arr))
}

// minSubArrayLen returns the minimal length of a contiguous subarray of which the sum ≥ target.
// If there is no such subarray, returns 0.
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minLen := math.MaxInt32
	sum := 0
	start := 0

	for end := 0; end < n; end++ {
		sum += nums[end]

		for sum >= target {
			windowLen := end - start + 1
			if windowLen < minLen {
				minLen = windowLen
			}
			sum -= nums[start]
			start++
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}

	return minLen
}
