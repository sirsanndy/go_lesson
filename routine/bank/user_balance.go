package bank

import (
	"fmt"
	"sync"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("User 1 Locked with balance : ", user1.Balance)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("User 2 Locked with balance : ", user2.Balance)
	user2.Change(amount)

	time.Sleep(1 * time.Second)
	user1.Unlock()
	user2.Unlock()
}
