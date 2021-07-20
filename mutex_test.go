package mylearngolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("counter = ", x)
}

type BankAccount struct {
	RWmutex sync.RWMutex
	Balance int
}

func (account *BankAccount) Addbalance(amount int) {
	account.RWmutex.Lock()
	account.Balance = account.Balance + amount
	account.RWmutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWmutex.RLock()
	balance := account.Balance
	account.RWmutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.Addbalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("total balance", account.GetBalance())
}

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
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "otong",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name:    "ucup",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 100000)

	time.Sleep(10 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance", user2.Balance)

}

// sync.waitgroup
// fitur yang bisa digunakan menunggu sebuah
func Asynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	// do anything here
	fmt.Println("hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go Asynchronous(group)

	}
	// now we no need run time.Sleep again
	group.Wait()
	fmt.Println("complete")

}
