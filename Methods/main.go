package main

import "fmt"

func main() {
	vinay := User{"Vinay", "vinay@gmail.com", true, 20}
	vinay.GetStatus()
	vinay.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	//do not change the original one copy of the user gets modified
	u.Email = "newEmail@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
