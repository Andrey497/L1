//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N int = 1

func PostInChannel(c chan interface{}, t <-chan time.Time) {
	for {
		select {
		case <-t:
			close(c)
			return
		default:
			random := rand.Int()
			c <- random
		}
	}
}

func main() {
	fmt.Println(time.Now())
	t := time.After(time.Duration(N) * time.Second)
	channel := make(chan interface{})
	go PostInChannel(channel, t)
	for data := range channel {
		fmt.Println(data)
	}
	fmt.Println(time.Now())

}
