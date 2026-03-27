package main
import "fmt"

func main(){
	var a = "Prachiti" //automatically convert as var a string
	fmt.Println(a)

	b := 10 //short form of var b = 10
	fmt.Println(b)

	var num int = 10 //commonly used, depends of 32/64-bit system-also int8,int16,int32,int64 avl
	fmt.Println(num)

	var pi float32 = 3.14
	fmt.Println(pi)

	var price float64 = 99.99 //more precise and used mostly
	fmt.Println(price)

	var isValid bool = true
	fmt.Println(isValid)

	var name string = "Kitey"
	fmt.Println(name)

	var (
		isTrue = true
		n = "john"
	)
	fmt.Println(isTrue, n)
}