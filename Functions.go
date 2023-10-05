package main

import "fmt"

//	func concat(s1 string, s2 string) string {
//		return s1 + s2
//	}

// func concat2(s1, s2 string, age int) string {
// 	return s1 + s2
// }

func getNames() (string, string) {
	return "Akshdeep", "Singh"
}

func main() {
	// fmt.Println(concat("Hello", "World"))
	// fmt.Println(concat2("Hello", "World", 21))

	//ignoring last name
	firstname, _ := getNames()
	fmt.Println("Your Name", firstname)

}
