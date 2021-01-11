package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw your money") // 직접 error 변수 정의

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit * amount on your account
// param으로 (ac Account)를 넘기면, copy가 넘어감
func (ac *Account) Deposit(amount int) {
	ac.balance += amount
}

// Balance of your account
func (ac Account) Balance() int {
	return ac.balance
}

// Withdraw your account
func (ac *Account) Withdraw(amount int) error {
	if ac.balance < amount {
		// 정의한 error 리턴
		return errNoMoney
		// return error.Error()
	}
	ac.balance -= amount
	return nil
}

// Change owner to new owner
func (ac *Account) ChangeOwner(newOwner string) {
	ac.owner = newOwner
}

// Owner of the account
func (ac Account) Owner() string {
	return ac.owner
}

func (ac Account) String() string {
	return fmt.Sprint("owner : ", ac.Owner(), "\nBalance : ", ac.Balance())
}
