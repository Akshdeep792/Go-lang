package main

import (
	"fmt"
)

func check_prime(num int) {
	ans := true
	for i := 2; i <= Sqrt(num); i++ {
		if num%2 == 0 {
			ans = false
			break
		}
	}

	if ans {
		fmt.Println("Prime Number")
	} else {
		fmt.Println("Not a prime number")
	}
}

func main() {

	Number := 9
	check_prime(Number)
}
