package main

import "fmt"

func main() {
	arr := []int{-100, -30, -3, 2, 4, 50, 90}

	var result []int
	for i := 0; i < len(arr)/2; i++ {
		a1 := arr[i] * arr[i]
		a2 := arr[len(arr)-1-i] * arr[len(arr)-1-i]
		if a1 > a2 {
			result = append(result, a1, a2)
			continue
		}
		result = append(result, a2, a1)
	}

	if len(arr)%2 != 0 {
		result = append(result, arr[len(arr)/2]*arr[len(arr)/2])
	}

	res := make([]int, 0, len(result))
	for i := 0; i < len(result); i++ {
		res = append(res, result[len(result)-1-i])
	}

	fmt.Println(res)
}
