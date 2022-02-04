//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func ReverceSentence(s string) string {
	sliceWord := strings.Fields(s)
	for i, j := 0, len(sliceWord)-1; i < j; i, j = i+1, j-1 {
		sliceWord[i], sliceWord[j] = sliceWord[j], sliceWord[i]
	}
	return strings.Join(sliceWord, " ")
}
func main() {
	fmt.Println(ReverceSentence("snow dog sun "))
}
