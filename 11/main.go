//Реализовать пересечение двух неупорядоченных множеств.
package main

import "fmt"

func IntersectionSet(set1 map[int]struct{}, set2 map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})
	for key, _ := range set1 {
		if _, ok := set2[key]; ok {
			result[key] = struct{}{}
		}
	}
	return result
}

func main() {
	array1 := []int{1, 10, 20, 3, 4, 7, 6}
	array2 := []int{2, 4, 20, 3, 4, 5, 1}
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})
	for _, value := range array1 {
		set1[value] = struct{}{}
	}
	for _, value := range array2 {
		set2[value] = struct{}{}
	}
	fmt.Println(IntersectionSet(set1, set2))
}
