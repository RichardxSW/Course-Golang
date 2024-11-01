package main

import (
	"fmt"
	// "time"
	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser = User{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthDate,
	// 	createdAt: time.Now(),
	// }

	// var appUser *User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	// appUser = &user.User{
	// 	FirstName: userFirstName,
	// 	LastName:  userLastName,
	// 	Birthdate: userBirthDate,
	// 	CreatedAt: time.Now(),
	// }

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "admin")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()
	// outputUserDetails(&appUser)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()  

	// ... do something awesome with that gathered data!
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}