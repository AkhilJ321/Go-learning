package main

import "fmt"

func main() {
	fmt.Println("array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "FFFFFF "

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"potato", "beans", "msfjison"}
	fmt.Println(vegList)
}
