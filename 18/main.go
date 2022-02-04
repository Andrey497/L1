//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Counter struct {
	Count int32
}

func (c *Counter) Increment() {
	atomic.AddInt32(&c.Count, 1)
}

func main() {
	counter := Counter{0}
	for i := 0; i < 10; i++ {
		go counter.Increment()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(counter.Count)
}
