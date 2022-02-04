//Реализовать паттерн «адаптер» на любом примере.
package main

import "L1/21/pkj"

func main() {
	client := pkj.Client{}
	wiredHeadPhones := pkj.WiredHeadPhones{}
	client.ConnectJackHeadPhones(&wiredHeadPhones)
	blHeadPhones := pkj.BluethoothHeadPhones{}
	adapter := pkj.BluethoothHeadphonesAdapter{
		&blHeadPhones,
	}
	adapter.InsertJack()

}
