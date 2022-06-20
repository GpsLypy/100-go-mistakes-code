package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu       sync.RWMutex
	balances map[string]float64
}

func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()
	c.balances[id] = balance
	c.mu.Unlock()
}

func (c *Cache) AverageBalance() float64 {
	c.mu.RLock()
	m := make(map[string]float64, len(c.balances))
	for k, v := range c.balances {
		m[k] = v
	}
	c.mu.RUnlock()

	sum := 0.
	for _, balance := range m {
		sum += balance
	}
	return sum / float64(len(m))
}

func main() {
	var (
		c  Cache
		wg sync.WaitGroup
	)
	c.balances = make(map[string]float64)

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 100; i++ {
			c.AddBalance(fmt.Sprintf("account%d", i), 100)
		}
	}()

	go func() {
		defer wg.Done()

		c.AverageBalance()
	}()

	wg.Wait()
}
