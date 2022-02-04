//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func NewtPoint(x int, y int) Point {
	return Point{x, y}
}

func (p Point) GetX() int {
	return p.x
}
func (p Point) GetY() int {
	return p.y
}

func Distance(point1 Point, point2 Point) float64 {
	x := point2.GetX() - point1.GetX()
	y := point2.GetY() - point1.GetY()
	return math.Sqrt(float64(x*x + y*y))
}
func main() {
	point1 := NewtPoint(5, 10)
	point2 := NewtPoint(1, 0)
	fmt.Println(Distance(point1, point2))

}
