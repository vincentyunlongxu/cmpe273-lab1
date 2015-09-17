package main

import ("testing"
		"time"
		"fmt"
		)

func sleepTest(t *testing.T){
	fmt.Println("1")
	testStart := time.Now()
	sleep(4)
	testLap := time.Now().Sub(testStart)
	
	if testLap < 4{
		t.Error("Error")
	}

}
