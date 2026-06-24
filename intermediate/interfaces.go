package intermediate

import (
	"fmt"
	"math"
	"reflect"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type rect1 struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func (r rect1) area() float64 {
	return r.height * r.width
}

func measure(g geometry) {
	fmt.Printf("Geometry: %v\n", g)
	fmt.Printf("Type: %s\n", getType(g))
	fmt.Printf("Area: %.2f\n", g.area())
	fmt.Printf("Perimeter: %.2f\n", g.perim())
}

func getType(i any) string {
	return reflect.TypeOf(i).Name()
}

// `any` is an alias for `interface{}` - either works here
func myPrinter(i ...any) {
	for idx, v := range i {
		fmt.Printf("arg[%d]: value=%v, type=%T\n", idx, v, v)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	// r1 := rect1{width: 5, height: 6}

	measure(r)
	measure(c)

	// rect1 does not implement geometry (missing perim)
	// so this would throw an error
	
	// measure(r1)

	myPrinter(1, "John", 45.9, true)
}
