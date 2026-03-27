package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter number : ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num>0{
		fmt.Println("Number is positive")
	} else if num<0{
		fmt.Println("Number is negative")
	} else{
		fmt.Println("Number is zero")
	}
}
