package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const fileName = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(fileName, 0)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		// panic("Can't continue, sorry.")
		// return
	}

	fmt.Println("Welcome", randomdata.FirstName(randomdata.Male))

	for {
		presentOptions()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		// if choice == 1 {
		// 	fmt.Println("Your balance is: ", balance)
		// } else if choice == 2 {
		// 	fmt.Print("Enter amount to deposit: ")
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount")
		// 		continue
		// 	}
		// 	balance += depositAmount
		// 	fmt.Println("Balance updated! New balance is: ", balance)

		// } else if choice == 3 {
		// 	var withdrawAmount float64
		// 	fmt.Print("Enter amount to withdraw: ")
		// 	fmt.Scan(&withdrawAmount)

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount")
		// 		continue
		// 	}

		// 	if withdrawAmount > balance {
		// 		fmt.Println("Insufficient balance")
		// 		continue
		// 	}
		// 	balance -= withdrawAmount
		// 	fmt.Println("Balance updated! New balance is: ", balance)

		// } else {
		// 	fmt.Println("Exiting...")
		// 	break
		// }

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Enter amount to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New balance is: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, fileName)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient balance")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New balance is: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, fileName)
		default:
			fmt.Println("Exiting...")
			fmt.Println("ThankYou")
			return
		}
	}
	// fmt.Println("ThankYou")
}
