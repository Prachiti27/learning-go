package main
import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	var name string
	fmt.Print("Enter your name : ")
	fmt.Scanln(&name)
	var age int
	fmt.Printf("Hi %s, enter your age : ",name)
	fmt.Scanln(&age)

	fmt.Print("Enter full name : ")
	var reader = bufio.NewReader(os.Stdin)
	var fname, _ = reader.ReadString('\n')
	fmt.Print("Full name is ",fname)
}