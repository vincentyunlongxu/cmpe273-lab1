package main

import "testing"

func TestMain(t *testing.T){
	r := Rectangle{length:5, width:3}
	var s Shape
	s = r
	if s.perimeter() != 16{
		t.Error("Expect 16, got ", s.perimeter())
	}

	c := Circle{pi:3.14, r:5}
	s = c
	if s.perimeter() != 31.400002{
		t.Error("Expect 31.400002, got ", s.perimeter())
	}
}
