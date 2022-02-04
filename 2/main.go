//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
package main

import "fmt"

var array2 = [5]int{2, 4, 6, 8, 10}

func SquareNumber(c chan int, number int) {
	square := number * number
	c <- square
}
func main() {

	c := make(chan int, 5) //канал размером 5 элиментов
	defer close(c)         //close channel
	for _, val := range array2 {
		//паралельное вычисление квадрата числа
		go SquareNumber(c, val)
	}
	//вывод значений из канала
	for range array2 {
		number := <-c
		fmt.Println(number)
	}

}
