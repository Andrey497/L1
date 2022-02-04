//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
package main

import "fmt"

type Human struct {
	firstName  string
	secondName string
	age        int
}

func (h *Human) ToString() string {
	return fmt.Sprintf("firstName: %s, secondName: %s, age: %d\n", h.firstName, h.secondName, h.age)
}

//1 Вариант встравивание
type Action struct {
	Human
}

//2 Вариант композиция
type Action2 struct {
	human Human
}

/*Главным отличием данных методов является что в первом случае Action наследует  все методы и мы можем вызывать их обращаясь напримую к обьекту Action,
а во втором случае нужно обратиться к свойсву Human, главным недостаткоми 1 го метода яввляется что если методы родителя и наследника будут иметь
одно и тоже имя возникнит ошибка, нельзя просто отделить один обьект от другого
*/
func main() {
	action := Action{Human{"Victor", "Victorov", 32}}
	fmt.Println(action.ToString())
	action2 := Action2{Human{"Victor", "Victorov", 32}}
	fmt.Println(action2.human)
}
