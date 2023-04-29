package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "javascritp"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("js shots for :", languages["JS"])

	delete(languages, "RB")
	fmt.Println("list of all languages: ", languages)

	// loops are interesting
	for key, value := range languages {
		fmt.Printf("For key %v , value is %v\n", key, value)
	}
}
