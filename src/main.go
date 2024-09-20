package main

import "fmt"

func main() {
	res := add(2, 3, 4)

	fmt.Printf("%d\n", res)
}

func add(nums ...int) int {
	sum := 0

	for _, val := range nums {
		sum += val
	}

	return sum
}
