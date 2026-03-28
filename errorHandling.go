package main

import "fmt"

func divide(a, b float64) (float64, error){
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a/b,nil
}

func main(){
	var a float64
	var b float64 
	fmt.Print("Enter a and b : ")
	// fmt.Scan(&a)
	// fmt.Scan(&b)

	fmt.Scan(&a, &b)

	result, err := divide(a, b)

	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	fmt.Println("Result:",result)
}