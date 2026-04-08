package main

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	longestOnes(nums, 2)
}

func longestOnes(nums []int, k int) {
	begin := 0
	windowState := 0
	result := 0

	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			windowState++
		}

		for windowState > k {
			result = max(result, end-begin+1)

			if nums[begin] == 0 {
				windowState -= 1
			}

			begin += 1 // shrink window
		}
	}

	return max(result, end-begin+1)
}
