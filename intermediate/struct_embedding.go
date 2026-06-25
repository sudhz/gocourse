package intermediate

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person // Embedded struct
	empId  string
	salary float64
}

func (p Person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f\n", e.name, e.empId, e.salary)
}

func main() {
	emp := Employee{
		Person: Person{
			name: "John",
			age:  30,
		},
		empId:  "E001",
		salary: 50000,
	}

	// Accessing embedded struct field directly (emp.person.name)
	fmt.Println("Name:", emp.name)
	fmt.Println("Age:", emp.age)
	fmt.Println("Emp ID:", emp.empId)
	fmt.Println("Salary:", emp.salary)

	// Overriden method
	emp.introduce()
}
