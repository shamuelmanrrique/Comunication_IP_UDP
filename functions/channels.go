package functions

import "fmt"

func food(text chan string) {
	fmt.Println("You are the best ", <-text)
}

// func main() {

// 	// localhost = '127.0.0.1'
// 	// ip = ''
// 	// var c chan int
// 	// fmt.Println(c)

// 	fmt.Println("Main is starting ")

// 	c := make(chan string)

// 	go food(c)

// 	c <- "shamuel"
// 	close(c)
// 	d, ok := <-c
// 	fmt.Println(ok)
// 	fmt.Println(d)
// 	fmt.Println("main() stopped")
// }
