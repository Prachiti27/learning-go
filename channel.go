//communication between go routines
package main

import (
	"fmt"
	"time"
	"math/rand"
)

//sending
func processNum(numChan chan int) {

	for num := range numChan {
		fmt.Println("Processing number", num)
		time.Sleep(time.Second)
	}
}

func sum(result chan int, num1 int, num2 int) {
	res := num1 + num2
	result <- res
}

//goroutine sync
func task(done chan bool) {
	defer func() { done <- true}()
	fmt.Println("processing")
	// done <- true
}

func main() {
	// messageChan := make(chan string)
	numChan := make(chan int)

	go processNum(numChan)

	for i:=0; i<5; i++ {
		numChan <- rand.Intn(100)
	}

	result := make(chan int)

	go sum(result, 4, 5)

	res := <- result
	fmt.Println(res)

	done := make(chan bool)

	go task(done)

	<- done //blocking

	//buffered channel
	email := make(chan string, 100)

	email <- "pk@gmail.com"
	email <- "sk@gmail.com"

	fmt.Println(<- email)
	fmt.Println(<- email)

	close(email) //closing channel
	//sending data
	// messageChan <- "ping"
	//receiving data
	// msg := <- messageChan
	// fmt.Println(msg)
}