package main

import "fmt"

func main() {
	loginCount := 2
	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if loginCount < 5 {
		result = "something else"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is evenb")

	} else {
		fmt.Println("odd number")
	}

	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}

}
