//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
package main

import "fmt"

func partition(arr *[]int, low int, high int) int {
	pivot := (*arr)[high] //опорный элемент
	wall := low - 1       //стена
	//проходим по всем жлиментам от стены до опорного не включая его
	for i := low; i <= high-1; i++ {
		//если элемент больше опорного оставляем его справа от стены
		//иначе меняем элемент с первым от стены и сдвигаем стену на 1 элемент вперед
		if (*arr)[i] < pivot {
			(*arr)[i], (*arr)[wall+1] = (*arr)[wall+1], (*arr)[i] //меняем элемент с первым от стены
			wall++                                                //сдвигаем стену на 1 элемент вперед
		}

	}
	(*arr)[wall+1], (*arr)[high] = (*arr)[high], (*arr)[wall+1] //меняем местами 1  элемент справа от стены и текущий
	return wall + 1
}

func QuickSort(arr *[]int, low int, high int) {
	if low < high { //условие выхода из сортировки
		pi := partition(arr, low, high)
		//рекурсивно вызываем сортировку слева и справа от текущего элемента
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println(arr)
	QuickSort(&arr, 0, len(arr)-1)
	fmt.Println(arr)
}
