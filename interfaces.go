package main

import "fmt"

type Shape interface{
	Area() float64
	Perimeter() float64
}

type Circle struct{
	radius float64
}

type Rectangle struct{
	length float64
	breadth float64
}

func (c Circle) Area() float64{
	return 3.14*c.radius*c.radius
}

func (r Rectangle) Area() float64{
	return r.length*r.breadth
}

func (c Circle) Perimeter() float64{
	return 2*3.14*c.radius
}

func (r Rectangle) Perimeter() float64{
	return 2*(r.length + r.breadth)
}

// func printArea(s Shape){
// 	area := s.Area()
// 	fmt.Println("The area is ",area)
// }

// func printPerimeter(s Shape){
// 	peri := s.Perimeter()
// 	fmt.Println("The perimeter is ",peri)
// }

func printShape(s Shape){
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n",s.Area(), s.Perimeter())
}

func main(){
	c := Circle{radius:2}
	r := Rectangle{length:2,breadth:2}
	// printShape(c)
	// printShape(r)

	shapes := []Shape{c, r}

	for _,s := range shapes{
		printShape(s)
	} 
}