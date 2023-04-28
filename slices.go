package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices tut")

	var fruitList = []string{"bbsdf", "ihfs", "oihfsd"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945

	highScores[2] = 465
	highScores[3] = 567
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}
