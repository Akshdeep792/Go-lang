package main

//	func concat(s1 string, s2 string) string {
//		return s1 + s2
//	}

// func concat2(s1, s2 string, age int) string {
// 	return s1 + s2
// }

// func getNames() (string, string) {
// 	return "Akshdeep", "Singh"
// }

// ----------------->NAMED RETURN VALUES
// func getDetails(age int) (tillAdult, tillDrink, tillRent int) {
// 	tillAdult = 18 - age
// 	if tillAdult < 0 {
// 		tillAdult = 0
// 	}

// 	tillDrink = 21 - age
// 	if tillDrink < 0 {
// 		tillDrink = 0
// 	}

// 	tillRent = 25 - age
// 	if tillRent < 0 {
// 		tillRent = 0
// 	}

// 	return
// }

func main() {
	// fmt.Println(concat("Hello", "World"))
	// fmt.Println(concat2("Hello", "World", 21))

	//ignoring last name --->ignoring return value
	// firstname, _ := getNames()
	// fmt.Println("Your Name", firstname)

	//-------------->Named return values
	// tillAdult, tillDrink, tillRent := getDetails(13)
	// fmt.Println("You are an adult in", tillAdult, "years")
	// fmt.Println("You can drink in", tillDrink, "years")
	// fmt.Println("You can rent a car in", tillRent, "years")
	// fmt.Println("====================================")

}
