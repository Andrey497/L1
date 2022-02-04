//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
package main

import "fmt"

func worker1(array *[]int, out chan int) {
	for _, data := range *array {
		out <- data
	}
	return
}
func worker2(in chan int, out chan int) {
	for data := range in {
		out <- data * data
	}
	close(in)
	return
}

func worker3(in chan int) {
	for data := range in {
		fmt.Println(data)
	}
	close(in)
	return
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go worker1(&array, channel1)
	go worker2(channel1, channel2)
	go worker3(channel2)
	fmt.Scanln()
}
