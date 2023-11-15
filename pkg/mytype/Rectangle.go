package mytype

type Rectangle struct { // Optional: Define Rectangle's own Data Structure
	width, height float64
}

func (r Rectangle) Area() float64 { // Must: Implement the first function of the interface
	return r.width * r.height
}

func (r Rectangle) Perim() float64 { // Must: Implement the second function of the interface
	return 2*r.width + 2*r.height
}
