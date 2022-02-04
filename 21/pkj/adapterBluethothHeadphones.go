package pkj

import (
	"fmt"
)

type BluethoothHeadphonesAdapter struct {
	*BluethoothHeadPhones
}

func (b *BluethoothHeadphonesAdapter) InsertJack() {
	fmt.Println("convert jack to Bluethooth")
	b.BluethoothConnect()
}
