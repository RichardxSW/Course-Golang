package user

import (
	"errors"
	"time"
	"fmt"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func New(userFirstName, userLastName, userBirthDate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("User must have a first name, last name, and birthdate")
	}
	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email: email,
		password: password,
		User: User{
			firstName: "SUPER",
			lastName: "ADMIN",
			birthdate: "------",
			createdAt: time.Now(),
		},
	}
}

func (u *User) OutputUserDetails() {
	fmt.Println("User: ", u.firstName, u.lastName, u.birthdate)
	fmt.Println(u.createdAt)
	fmt.Println("Age: ", time.Since(u.createdAt).Hours()/24/365)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}