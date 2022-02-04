//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

type BigInt big.Int

func NewBigInt(number int64) *BigInt {
	return (*BigInt)(big.NewInt(number))
}
func (b *BigInt) Add(b2 *BigInt) *BigInt {
	result := new(big.Int).Add((*big.Int)(b), (*big.Int)(b2))
	return (*BigInt)(result)
}
func (b *BigInt) Mul(b2 *BigInt) *BigInt {
	result := new(big.Int).Mul((*big.Int)(b), (*big.Int)(b2))
	return (*BigInt)(result)
}

func (b *BigInt) Div(b2 *BigInt) *BigInt {
	result := new(big.Int).Div((*big.Int)(b), (*big.Int)(b2))
	return (*BigInt)(result)
}
func (b *BigInt) Sub(b2 *BigInt) *BigInt {
	result := new(big.Int).Sub((*big.Int)(b), (*big.Int)(b2))
	return (*BigInt)(result)
}

func (b *BigInt) Print() {
	fmt.Println((*big.Int)(b))
}

func main() {
	a := NewBigInt(400000000)
	b := NewBigInt(200000000)
	a.Add(b).Print()
	a.Sub(b).Print()
	a.Mul(b).Print()
	a.Div(b).Print()
}
