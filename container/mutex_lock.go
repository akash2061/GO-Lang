package main

import (
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int, i int) {
	pl("Attempt:", i)
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		pl("Not Enough Money in Account")
	} else {
		pl(v, "Withdrawn : Balance:", a.balance)
		a.balance -= v
	}
}

func main() {
	var acc Account
	acc.balance = 150

	pl("Balance:", acc.GetBalance())
	for i := 0; i < 12; i++ {
		// pl("Attempt:", i)
		go acc.Withdraw(15, i)
	}

	time.Sleep(2 * time.Second)

	for i := 0; i < 5; i++ {
		pl("Closures:")
		val := func() {
			i += 1
		}
		val()
		pl("Closure of I:", i)
	}
}
