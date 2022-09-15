package main

import (
	"fmt"
	"os"
)

type user struct {
	userid string
	pwd    string
}

var option string

func loginfnc() {
	var operation string

	fmt.Println("Login")
	fmt.Printf("Enter the user name:")
	fmt.Scan(&curUser.userid)
	fmt.Printf("Enter the password:")
	fmt.Scan(&curUser.pwd)
	if curUser.userid == "Test" && curUser.pwd == "test" {
	opt:
		fmt.Println("Please select the option to proceed:")
		fmt.Println("\nl -> WithDraw\n2 -> Deposit\n5 -> Balance\n4 -> Quit")
		fmt.Scan(&operation)

		switch operation {
		case "1":

			fmt.Println("Please collect the cash")
			goto opt
		case "2":
			fmt.Println("Amount deposited")
			goto opt
		case "3":
			fmt.Println("Account Balance is XXXXX")
			goto opt
		case "4":
			fmt.Println("Thank you for banking with us")
			os.Exit(4)
		}

	} else {
		main()
	}
}
func createAcct() {
	newUser := user{}
	fmt.Println("Enter the user name")
	fmt.Scan(&newUser.userid)
	fmt.Println("Enter the password")
	fmt.Scan(&newUser.pwd)
	fmt.Println("Acct created")
	main()
}

var curUser user

func main() {
	fmt.Println("Hi! Welcome to Mr. Rohit ATM Machine!\n\nPlease select an option from the menu below:")
	fmt.Println("\ni -> Login\nc -> Create New Account\nq -> Quit")
	fmt.Scan(&option)
	switch option {
	case "i":
		loginfnc()
	case "c":
		createAcct()
	case "q":
		fmt.Println("Thank you for Banking with us..")
		os.Exit(4)
	}

}
