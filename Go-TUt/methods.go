package main

import "fmt"

func main() {
	// no ingeritance in golang. no super or parent
	hitesh := User{"Hitseh", "skldf@gldng.gfd", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("details are : %+v\n", hitesh)
	fmt.Printf("Name : %v and email %v.\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)
}
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email fo this user is:", u.Email)
}
