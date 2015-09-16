ackage main
	
import "fmt"

func main() {
	var num int
	var result int
	fmt.Print("please Enter a Number: ")
	fmt.Scan(&num)
	result = fibonacci(num)
	fmt.Println("Your Answer is: ", result)

}

func fibonacci(num int) int {
	if num == 0{
		return 0
	}else if num == 1 {
		return 1
	}else{
		return fibonacci(num - 1) + fibonacci(num - 2)
	}
}
