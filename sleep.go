package main

import ("time"
		"fmt")

func main() {
	c := make(chan string, 1)
	fmt.Println("1")
	go func () {
		sleep(4)
		fmt.Println("2")
		c <- "2"
	}()
	fmt.Println("3")
	select{
	case result := <- c:
		fmt.Println(result)
	}
}

func sleep(num int) {
	<-time.After(time.Second * time.Duration(num))
}
