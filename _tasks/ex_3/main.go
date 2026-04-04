package main

import "fmt"

// tickets := []City{
// 	{
// 		from: "Moscow",
// 		to:   "Saint Petersburg",
// 	},
// 	{
// 		from: "Saint Petersburg",
// 		to:   "Kirov",
// 	},
// 	{
// 		from: "Kirov",
// 		to:   "Almaty",
// 	},
// 	{
// 		from: "Almaty",
// 		to:   "Tbilisi",
// 	},
// }

type Flight struct {
	from string
	to   string
}

func main() {
	println("start")
	tickets := []Flight{
		{
			from: "Almaty",
			to:   "Tbilisi",
		},
		{
			from: "Moscow",
			to:   "Saint Petersburg",
		},
		{
			from: "Kirov",
			to:   "Almaty",
		},
		{
			from: "Saint Petersburg",
			to:   "Kirov",
		},
	}

	mm := make(map[string]Flight)
	destinations := make(map[string]bool)

	for _, v := range tickets {
		mm[v.from] = v
		destinations[v.to] = true
	}

	result := make([]Flight, len(tickets))
	for _, v := range tickets {
		if !destinations[v.from] {
			result[0] = v
			break
		}
	}

	for i := 1; i < len(tickets); i++ {
		result[i] = mm[result[i-1].to]
	}

	fmt.Println(result)
}
