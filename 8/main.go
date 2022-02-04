//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0
package main

import "fmt"

func SetBit(number int64, index int, value bool) int64 {
	if value == true {
		return number | (1 << index) //nubmer | 1 and n(0)
	}
	return number &^ (1 << index) //nubmer & not(1 and n(0))

}

func main() {
	var x int64
	x = 100
	fmt.Println(SetBit(x, 1, true))

}
