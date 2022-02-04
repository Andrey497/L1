//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func AddWorcker(c chan interface{}, id int) {

	for data := range c {
		fmt.Println(data)
		fmt.Println(id)
	}
	fmt.Printf("Воркер %d завершил работу\n", id)
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	c := make(chan interface{})

	for i := 0; i < 5; i++ {
		go AddWorcker(c, i)
	}
	for {
		select {
		case <-sigChan:
			close(c)
			time.Sleep(time.Millisecond)
			os.Exit(0)

		default:
			random := rand.Int()
			c <- random
		}
	}

}
