package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}
	var v uint64

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)
}
