package routine

import (
	"fmt"
	"go-routine/bank"
	"testing"
	"time"
)

func TestDeadlock(t *testing.T) {
	user1 := bank.UserBalance{
		Name:    "Sandy",
		Balance: 1500000,
	}

	user2 := bank.UserBalance{
		Name:    "Addy",
		Balance: 3000000,
	}

	go bank.Transfer(&user1, &user2, 100000)
	go bank.Transfer(&user2, &user1, 300000)

	time.Sleep(3 * time.Second)
	fmt.Println("User", user1.Name+", Balance", user1.Balance)
	fmt.Println("User", user2.Name+", Balance", user2.Balance)
}
