package intermediate

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("The stack is empty")
		return
	}

	fmt.Print("Stack elements: ")
	for i, element := range s.elements {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(element)
	}
	fmt.Println()
}

func main() {
	x1, y1 := 1, 2
	x1, y1 = swap(x1, y1)
	fmt.Printf("Swapped ints:   %d, %d\n", x1, y1)

	x2, y2 := "John", "Jane"
	x2, y2 = swap(x2, y2)
	fmt.Printf("Swapped strings: %s, %s\n", x2, y2)

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)

	intStack.printAll()

	element, ok := intStack.pop()
	fmt.Printf("Popped: %v (ok: %v)\n", element, ok)
	element, ok = intStack.pop()
	fmt.Printf("Popped: %v (ok: %v)\n", element, ok)
	element, ok = intStack.pop()
	fmt.Printf("Popped: %v (ok: %v)\n", element, ok)

	fmt.Printf("Stack empty? %v\n", intStack.isEmpty())

	stringStack := Stack[string]{}
	stringStack.push("Hello")
	stringStack.push("World")
	stringStack.push("John")

	stringStack.printAll()

	sElement, sOk := stringStack.pop()
	fmt.Printf("Popped: %v (ok: %v)\n", sElement, sOk)
	sElement, sOk = stringStack.pop()
	fmt.Printf("Popped: %v (ok: %v)\n", sElement, sOk)
	sElement, sOk = stringStack.pop()
	fmt.Printf("Popped: %v (ok: %v)\n", sElement, sOk)

	fmt.Printf("Stack empty? %v\n", stringStack.isEmpty())
}
