package main

import (
	"fmt"
	"time"
)

func Sleep(s uint) {
	timer := time.After(time.Duration(s) * time.Millisecond)
	select {
	case <-timer:
		return
	}
}

func main() {
	fmt.Println(time.Now())
	Sleep(10000)
	fmt.Println(time.Now())
}
