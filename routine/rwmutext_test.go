package routine

import (
	"fmt"
	"go-routine/bank"
	"testing"
	"time"
)

func TestRWMutexBankAccount(t *testing.T) {
	account := bank.BankAccount{}
	for i := 0; i < 1000; i++ {
		go func() {
			account.AddBalance(1)
			fmt.Println(account.GetBalance())
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}
