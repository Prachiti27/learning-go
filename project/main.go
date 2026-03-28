package main

import (
	"fmt"
	"project/utils"
)

func main(){
	result := utils.Add(10, 20)
	fmt.Println("Sum: ",result)

	result = utils.Sub(10, 20)
	fmt.Println("Subtraction: ", result)

	result = utils.Mul(10, 20)
	fmt.Println("Multiplication: ", result)
}