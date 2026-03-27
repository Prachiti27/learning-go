package main
import "fmt"

func sum(a int, b int) int {
	return a + b
}

func fact(n int) int {
	var f int = 1
	for i:=1; i<=n; i++ {
		f *= i 
	}
	return f
}

func fib(n int) int {
	if n<=1 {
		return n
	}
	a,b := 0, 1
	for i:=2;i<=n;i++ {
		c := a+b
		a = b
		b = c
	}
	return b
}

func main(){
	fmt.Println(sum(2,3))
	fmt.Println(fact(5))
	fmt.Println(fib(5))
}