//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
package main

import (
	"fmt"
	"reflect"
)

func Type_1(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	case chan interface{}:
		fmt.Println("channel", v)
	default:
		fmt.Println("other type", v)
	}
}
func Type_2(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
}

func main() {
	var a int
	var testElement interface{}
	testElement = a
	Type_1(testElement)
	Type_2(testElement)
	var b string
	testElement = b
	Type_1(testElement)
	Type_2(testElement)
	var c bool
	testElement = c
	Type_1(testElement)
	Type_2(testElement)
	var d chan interface{}
	testElement = d
	Type_1(testElement)
	Type_2(testElement)
}
