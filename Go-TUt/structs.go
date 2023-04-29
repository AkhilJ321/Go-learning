package main

import "fmt"

func main() {
	// no ingeritance in golang. no super or parent
	hitesh := User{"Hitseh", "skldf@gldng.gfd", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("details are : %+v\n", hitesh)
	fmt.Printf("Name : %v and email %v.", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
