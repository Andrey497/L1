//Реализовать все возможные способы остановки выполнения горутины.
package main

import (
	"context"
	"fmt"
	"time"
)

//1 способ создать канал при подачи сообщения в который горутина закрывается
func worker1(c chan interface{}, quit <-chan struct{}) {
	for {
		select {
		case c <- time.Now():
		case <-quit:
			close(c)
			fmt.Println("close Worker 1")
			return
		}
	}
}

//2 способ использование контекста
func worker2(ctx context.Context, c chan interface{}) {
	for {
		select {
		case c <- time.Now():
		case <-ctx.Done():
			close(c)
			fmt.Println("close Worker 2")
			return
		}
	}
}

func main() {
	quitChannel := make(chan struct{})
	dataChannel := make(chan interface{})
	ctx, cancel := context.WithCancel(context.Background())

	//go worker1(dataChannel, quitChannel)
	go worker2(ctx, dataChannel)

	//close worker 1
	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		quitChannel <- struct{}{}
	}()

	//close worker 2
	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		cancel()
	}()
	for data := range dataChannel {
		fmt.Println(data)
	}
	fmt.Scanln()
}
