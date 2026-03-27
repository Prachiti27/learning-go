package main

import "fmt"

type Student struct {
	name string
	marks int
}

func (s Student) IsPassed() {
	if s.marks > 50{
		fmt.Printf("%s is passed\n",s.name)
	} else{
		fmt.Printf("%s is failed\n", s.name)
	}
}

func (s Student) UpdateMarks(newMarks int){
	s.marks = newMarks
}

func main(){
	s1 := Student{"Prachiti", 80}
	s1.IsPassed()
	s2 := Student{"Shreya", 30}
	s2.IsPassed()

	s2.UpdateMarks(60)
	s2.IsPassed()
}