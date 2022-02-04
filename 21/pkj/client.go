package pkj

import (
	"fmt"
)

type Client struct {
}

func (c *Client) ConnectJackHeadPhones(headphones Headphones) {
	fmt.Println("Client   connected  via jack headphones.")
	headphones.InsertJack()
}
