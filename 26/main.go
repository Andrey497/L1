//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
// aabcd — false
package main

import (
	"fmt"
	"strings"
)

func UniqueCheck(s string) bool {
	set := make(map[rune]struct{})
	stringLower := strings.ToLower(s)
	runes := []rune(stringLower)
	for _, val := range runes {
		if _, ok := set[val]; ok {
			return false
		}
		set[val] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(UniqueCheck("abcd"))
	fmt.Println(UniqueCheck("abCdefAaf"))
	fmt.Println(UniqueCheck("aabcd"))
}
