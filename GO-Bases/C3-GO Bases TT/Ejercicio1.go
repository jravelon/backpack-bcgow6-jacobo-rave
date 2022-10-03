package main

import "fmt"

type user struct {
	name     string
	age      int
	mail     string
	password string
}

func (u *user) setName(name string) {
	u.name = name
}

func (u *user) setAge(age int) {
	u.age = age
}

func (u *user) setMail(mail string) {
	u.mail = mail
}

func (u *user) setPassword(password string) {
	u.password = password
}

func (u *user) printInfo() {
	fmt.Printf("Name: %v\nAge: %v\nMail: %v\nPassword: %v\n", u.name, u.age, u.mail, u.password)
}

func main() {
	u1 := user{"Jacobo", 20, "jevar", "1323"}
	u1.printInfo()
	u1.setName("Juan")
	u1.printInfo()
}
