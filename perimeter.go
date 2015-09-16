package main

import "fmt"

type Shape interface{
	perimeter() float32
}

type Rectangle struct{
	length float32
	width float32
}

type Circle struct{
	pi float32
	r float32
}

func (r Rectangle) perimeter() float32 {
	return 2*(r.length+r.width)
}

func (c Circle) perimeter() float32 {
	return 2*c.pi*c.r
}

func main() {
	r := Rectangle{length:5, width:3}
	var s Shape
	s = r
	fmt.Println("Rectangle's perimeter is: ", s.perimeter())

	c := Circle{pi:3.14, r:5}
	s = c
	fmt.Println("Circle's perimeter is: ", s.perimeter())
	fmt.Println()

}
