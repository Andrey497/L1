package pkj

import "fmt"

type BluethoothHeadPhones struct{}

func (w *BluethoothHeadPhones) BluethoothConnect() {
	fmt.Println("Bluethooth connect")
}
