package bankSem2

import (
	"sync"
)

var (
	mu      sync.RWMutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func deposit(amount int) {
	balance += amount
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}
