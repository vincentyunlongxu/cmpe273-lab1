package main

import "testing"

func TestFibonacci(t *testing.T){
	var result int
	result = fibonacci(10)
	if result != 55{
		t.Error("Expected 55, got ", result)
	}
}
