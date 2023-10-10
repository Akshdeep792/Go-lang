package main

import (
	"fmt"
)

func main() {
	total_cost := 0.0
	num_of_msg := 34
	for i := 0; i < num_of_msg; i++ {
		total_cost += 1.0 + (0.01 * float64(i))
	}

	fmt.Println("%f", total_cost)
}
