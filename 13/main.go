//Задача:Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func SwapInt(x1 *int, x2 *int) {
	*x1 = *x1 + *x2
	*x2 = *x1 - *x2
	*x1 = *x1 - *x2
}

func main() {
	a := -100
	b := 101
	SwapInt(&a, &b)
	//2 вариант
	//a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}
