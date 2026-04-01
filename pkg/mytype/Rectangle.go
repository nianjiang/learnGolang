package mytype

import "fmt"

type Rectangle struct { // Optional: Define Rectangle's own Data Structure
	Width, Height float64
}

func (r Rectangle) Area() float64 { // Must: Implement the first function of the interface
	return r.Width * r.Height
}

func (r Rectangle) Perim() float64 { // Must: Implement the second function of the interface
	return 2*r.Width + 2*r.Height
}

func Test() {
	r := Rectangle{Width: 3, Height: 4}
	fmt.Println(r)
}
