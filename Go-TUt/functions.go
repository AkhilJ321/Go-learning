package main

import "fmt"

func main() {
	greeter()
	result := adder(3, 5)
	fmt.Println("result is:", result)

	proRes, mymsg := proAdder(2, 5, 8, 7)
	fmt.Println(proRes, mymsg)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "fi prosdlfn"
}

func greeterTwo() {
	fmt.Println("ANother method")
}
func greeter() {
	fmt.Println("namaste from golang")
}
