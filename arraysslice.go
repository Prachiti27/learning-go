package main
import "fmt"

func main(){
	s := []int{1,2,3,56} //slice

	min := s[0]
	max := s[0]

	for i:=0; i<4; i++{
		if s[i]>max{
			max = s[i]
		}
		if s[i]<min{
			min = s[i]
		}
	}
	fmt.Printf("Minimum is %d and Maximum is %d.\n",min,max)

	i := 0
	j := 3

	for i<j{
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	for i:=0; i<4; i++{
		fmt.Println(s[i])
	}
}