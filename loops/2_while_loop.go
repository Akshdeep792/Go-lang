package main

import (
	"fmt"
)

func main() {
	total_cost := 0.0
	num_of_msg := 0
	for num_of_msg < 5 {
		total_cost += 1.0 + (0.01 * float64(num_of_msg))
		num_of_msg++
	}

	fmt.Println("%f", total_cost)
}
