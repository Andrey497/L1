//Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
	"time"
)

type SyncMap struct {
	sync.RWMutex
	dict map[int]int
}

const N int = 3

func worker(s *SyncMap, idWorker int) {
	for i := 0; i < 20; i++ {
		s.Lock()
		s.dict[i] = idWorker
		s.Unlock()
		fmt.Printf("Воркер %v написал по ключу %v значение %v\n", idWorker, i, idWorker)
	}
}

func main() {
	cache := SyncMap{
		RWMutex: sync.RWMutex{},
		dict:    map[int]int{},
	}
	for i := 0; i < N; i++ {
		go worker(&cache, i)
	}
	time.Sleep(time.Millisecond)
	cache.RLock() //
	for key, value := range cache.dict {
		fmt.Printf("ключ: %v\tзначение: %v\n", key, value)
	}
	cache.RUnlock()
	fmt.Scanln()
}
