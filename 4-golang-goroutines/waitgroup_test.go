package main_test

import (
	"fmt"
	"sync"
	"testing"
)

type BankAccount struct {
	sync.RWMutex
	Balance int
}

func (b *BankAccount) Lock() {
	b.RWMutex.Lock()
}

func (b *BankAccount) Unlock() {
	b.RWMutex.Unlock()
}

func (b *BankAccount) RLock() {
	b.RWMutex.RLock()
}

func (b *BankAccount) RUnlock() {
	b.RWMutex.RUnlock()
}

func (b *BankAccount) AddBalance(wg *sync.WaitGroup, amount int) {
	defer b.Unlock()
	b.Lock()
	b.Balance += amount
	wg.Done()
}

func (b *BankAccount) GetBalance() int {
	b.RLock()
	balance := b.Balance
	b.RUnlock()

	return balance
}

func TestWaitGroup(t *testing.T) {
	account := BankAccount{}
	
	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go account.AddBalance(wg, 1)
	}

	wg.Wait()
	fmt.Println("Total Balance:", account.GetBalance())
}