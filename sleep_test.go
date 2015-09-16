package main

import ("testing"
		"time"
		"fmt")

func sleepTest(t *testing.T){

	num := 4

	c := make(chan int, 1)
	fmt.Println("1")
	go func () {
		testStart := time.Now().Second()
		sleep(num)
		testEnd := time.Now().Second()
		fmt.Println("2")
		c <- (testEnd - testStart)
	}()
	fmt.Println("3")

	result := <- c

	if result < num{
		t.Error("Error")
	}

}
