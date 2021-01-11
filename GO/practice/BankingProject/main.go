package main

import (
	"fmt"

	"./accounts"
)

func main() {
	account := accounts.NewAccount("lee")
	fmt.Println("current status : ", account)

	account.Deposit(20)
	fmt.Println("after deposit money", account.Balance())

	err := account.Withdraw(30)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("after withdraw money", account.Balance())

	account.ChangeOwner("Ryan")

	fmt.Println("New Owner : ", account.Owner(), account.Balance())
	fmt.Println("String : ", account)
}
