package mylearngolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//Pool's purpose is to cache allocated but unused items for later reuse, relieving pressure on the garbage collector. That is, it makes it easy to build efficient, thread-safe free lists. However, it is not suitable for all free lists.
func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "empty"
		},
	}
	var group sync.WaitGroup
	pool.Put("maiing")
	pool.Put("a Pool is safe for use by multiple goroutines simultaneously ")
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()

	}
	// time.Sleep(5 * time.Second)
	group.Wait()
	fmt.Println("selesai")
}
