package main

import "fmt"

func main() {
	days := []string{"sunday", "tuesdAY", "wednesday", "friday", "saturday"}

	fmt.Println(days)
	// conventional loop
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i:= range days{
	// 	fmt.Println(days[i])
	// }
	for _, day := range days {
		fmt.Printf("index is and value is %v", day)
	}

}
