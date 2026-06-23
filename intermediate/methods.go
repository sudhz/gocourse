package intermediate

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

type MyInt int

// Method on a user defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt type"
}

func main() {
	rect := Rectangle{
		length: 10,
		width:  9,
	}

	area := rect.Area()
	fmt.Printf("Area of rectangle with length %.2f and width %.2f: %.2f\n", rect.length, rect.width, area)

	scaleFactor := 2.0
	rect.Scale(scaleFactor)
	scaledArea := rect.Area()
	fmt.Printf("After scaling by %.2f - length: %.2f, width: %.2f, area: %.2f\n", scaleFactor, rect.length, rect.width, scaledArea)

	num := MyInt(-5)
	num1 := MyInt(9)

	fmt.Printf("num = %d, IsPositive: %t\n", num, num.IsPositive())
	fmt.Printf("num1 = %d, IsPositive: %t\n", num1, num1.IsPositive())
	fmt.Println(num.welcomeMessage())

	s := Shape{
		Rectangle: Rectangle{
			length: 10,
			width:  9,
		},
	}

	fmt.Printf("Shape (embedded Rectangle) area: length %.2f * width %.2f = %.2f\n", s.length, s.width, s.Area())
	fmt.Printf("Accessing embedded field directly - s.Rectangle.Area(): %.2f\n", s.Rectangle.Area())
}
