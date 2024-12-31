package main_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserAccount struct {
	sync.Mutex
	UserName string
	Balance int
}

func (u *UserAccount) Lock() {
	u.Mutex.Lock()
}

func (u *UserAccount) Unlock() {
	u.Mutex.Unlock()
}

func (u *UserAccount) change(amount int) {
	u.Balance += amount
}

func Transfer(user1 *UserAccount, user2 *UserAccount, amount int) {
	user1.Lock()
	fmt.Println("Lock User1:", user1.UserName)
	user1.change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User2:", user2.UserName)
	user2.change(amount)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserAccount{
		UserName: "Nabil",
		Balance: 100000,
	}

	user2 := UserAccount{
		UserName: "Hamzah",
		Balance: 100000,
	}

	go Transfer(&user1, &user2, 10000)
	go Transfer(&user2, &user1, 20000)

	time.Sleep(5 * time.Second)

	fmt.Println("Saldo", user1.UserName + "Sisa Balance", user1.Balance)
	fmt.Println("Saldo", user2.UserName + "Sisa Balance", user2.Balance)
} 