//Удалить i-ый элемент из слайса.
package main

import "fmt"

func RemoveElement(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}
func main() {

	s := []int{1, 2, 3, 4, 5}
	s = RemoveElement(s, 2)
	fmt.Print(s)
}
