//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
package main

import "fmt"

func SquareNumber(c chan int, number int) {
	square := number * number
	c <- square
}
func main() {
	sum := 0
	var array3 = [...]int{2, 4, 6, 8, 10}
	c := make(chan int, 5) //канал размером 5 элиментов
	defer close(c)         //close channel
	for _, val := range array3 {
		//паралельное вычисление квадрата числа
		go SquareNumber(c, val)
	}
	//sum  square numbers
	for range array3 {
		sum += <-c
	}
	fmt.Println(sum)
}
