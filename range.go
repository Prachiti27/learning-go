package main
import "fmt"

func main(){
	var n int
	fmt.Print("Enter number of elements in array : ")
	fmt.Scan(&n)

	arr := make([]int,n)
	fmt.Println("Enter elements in array : ")
	for i:=0; i<n;i++{
		fmt.Scan(&arr[i])
	}
	
	for i,v:= range arr{
		fmt.Println(i,v)
	}

	sum := 0
	for _,v := range arr{
		sum += v
	}
	fmt.Printf("Sum using range is %d.",sum)

	m := map[string]int{
		"a":1,
		"b":2,
		"c":3
	}
	for k,v := range m{
		fmt.Println(k,v)
	}
	//delete index at i
	i := 3
	arr = append(arr[:i], arr[i+1:]...)

	x := 3
	arr = append(arr[:i], append([]int{x}, arr[i:]...)...)

	//reverse ik so no need to code here
}