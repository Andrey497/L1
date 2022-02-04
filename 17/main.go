//Реализовать бинарный поиск встроенными методами языка.
package main

import "fmt"

func binarySearch(array *[]int, element int, left int, right int) int {
	if right >= left {
		mid := (left + right) / 2
		if (*array)[mid] == element {
			return mid
		}
		if (*array)[mid] < element {
			return binarySearch(array, element, mid+1, right)
		}

		return binarySearch(array, element, left, right+1)

	}
	return -1
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(&array, 8, 0, len(array)-1))

}
