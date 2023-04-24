package main

import (
	"fmt"
	"sync"
)

type Counter struct { //структура счетчика
	count int
	mu    sync.Mutex
}

func main() {
	wg := new(sync.WaitGroup)
	c := Counter{}
	for i := 0; i < 1000; i++ {
		// Запускаем 1000 экземпляров горутины, увеличивающей счетчик на 1
		wg.Add(1)
		go func(wg *sync.WaitGroup, c *Counter) {
			defer wg.Done()
			c.mu.Lock()
			c.count++
			c.mu.Unlock()
		}(wg, &c)
	}

	wg.Wait()
	fmt.Println(c.count)
}
