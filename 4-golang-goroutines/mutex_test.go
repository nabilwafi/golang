package main_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	counter := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func () {
			for i := 0; i < 100; i++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter:", counter)
}

type EWallet struct {
	RWMutex sync.RWMutex
	Balance int
}

func (wallet *EWallet) addBalance(amount int) {
	wallet.RWMutex.Lock()
	wallet.Balance += amount
	wallet.RWMutex.Unlock()
}

func (wallet *EWallet) getBalance() int {
	wallet.RWMutex.RLock()
	balance := wallet.Balance
	wallet.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := EWallet{}

	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				account.addBalance(1)
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Total Balance: ", account.getBalance())
}