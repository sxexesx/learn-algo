package main

import "fmt"

func main() {
	k := 4
	nums := []int{1, 12, -5, -6, 50, 3}

	begin := 0
	windowState := 0 // состояние рамки
	result := 0

	for end := 0; end < len(nums); end++ {
		windowState += nums[end]

		if end-begin+1 == k {
			result = max(result, windowState)
			windowState -= nums[begin]
			begin += 1
		}
	}

	fmt.Println(result / 4)

}
