package pkj

import "fmt"

type WiredHeadPhones struct {
}

func (w *WiredHeadPhones) InsertJack() {
	fmt.Println("Insert to Jack connect")
}
