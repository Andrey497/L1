//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
/*
var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}
func main() {
someFunc()
}
*/
package main

import (
	"fmt"
)

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "Б"
	}
	return v
}

var justString string
var justString2 string

func someFunc() {

	v := createHugeString(1 << 10)
	//проблема заключается в том, что данный код выведет последние 100 байт, а не 100 символов
	//например в данном случае вывел 50 символов

	justString = v[:100]
	//коректный код
	justString2 = string(([]rune(v))[:100])

	fmt.Println("v:", v)
	fmt.Println("justString: ", justString)
	fmt.Println("len(v)", len(v))
	fmt.Println("len(justString)", len(justString))
	fmt.Println("CountRun(justString)", len([]rune(justString)))
	fmt.Println("len(justString2)", len(justString2))
	fmt.Println("CountRun(justString2)", len([]rune(justString2)))
}

func main() {
	someFunc()

}
