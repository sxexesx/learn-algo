package main

import "fmt"

func main() {
	arr := []int{1, 3, 4}
	target := 7

	fmt.Println(findSum(arr, target))
}

func findSum(arr []int, target int) []int {
	mm := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		a := target - arr[i]
		if j, ok := mm[a]; ok {
			return []int{i, j}
		}
		mm[arr[i]] = i
	}
	return []int{}
}
