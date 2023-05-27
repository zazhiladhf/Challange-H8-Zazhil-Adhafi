package entity

import "fmt"

type User struct {
	ID       int
	UserName string
	Email    string
	Password string
}

func (u *User) PrintDetail() {
	fmt.Println("ID : ", u.ID)
	fmt.Println("Username : ", u.UserName)
	fmt.Println("Email : ", u.Email)
	fmt.Println("Password : ", u.Password)
}
