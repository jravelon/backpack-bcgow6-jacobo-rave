package main

import "fmt"

type User struct {
	Name     string
	Surname  string
	Age      int
	Email    string
	Password string
}

func (u *User) ChangeName(newName, newSurname string) {
	u.Name = newName
	u.Surname = newSurname
}

func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}

func (u *User) ChangeEmail(newEmail string) {
	u.Email = newEmail
}

func (u *User) ChangePassword(newPassword string) {
	u.Password = newPassword
}

func (u *User) PrintInfo() {
	fmt.Printf("Name: %v\nSurname: %v\nAge: %v\nMail: %v\nPassword: %v\n", u.Name, u.Surname, u.Age, u.Email, u.Password)
}

func main() {
	user := User{
		"Jacobo", "Rave", 19, "j@gmail.com", "hola",
	}
	user.PrintInfo()
	user.ChangeName("Julio", "Lazio")
	user.ChangeAge(20)
	user.ChangeEmail("juan@gmail.com")
	user.ChangePassword("pelicano")
	user.PrintInfo()
}
