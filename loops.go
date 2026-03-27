package main

import "fmt"

func main(){
	var n int
	fmt.Print("Enter value of n : ")
	fmt.Scanln(&n)
	
	fmt.Print("Numbers: ")
	for i:=1; i<=n; i++{
		fmt.Printf("%d ",i)
	}
	fmt.Println()

	var sum int = 0

	var i int = 0
	for i<=n {
		sum += i 
		i++
	}
	fmt.Printf("Sum of numbers is : %d\n",sum)
	for i:=1; i<=10; i++ {
		fmt.Printf("%d * %d = %d\n",i,n,i*n)
	}
}