package intermediate

import "fmt"

func main() {
	// Printing functions
	fmt.Print("Hello ")
	fmt.Print("World!")
	fmt.Print(12, 456)

	fmt.Println("Hello ")
	fmt.Println("World!")
	fmt.Println(12, 456)

	name := "John"
	age := 25

	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// Formatting functions
	s := fmt.Sprint("Hello", "World!", 123, 456)
	fmt.Print(s)

	s = fmt.Sprintln("Hello", "World!", 123, 456)
	fmt.Print(s)

	sf := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(sf)

	// Scanning functions
	var scan_name string
	var scan_age int

	fmt.Print("Enter your name and age: ")
	fmt.Scanln(&scan_name, &scan_age)
	fmt.Printf("Name: %s, Age: %d\n", scan_name, scan_age)

	var scan_city string
	var scan_zip int

	fmt.Print("Enter your city and zip code: ")
	fmt.Scan(&scan_city, &scan_zip)
	fmt.Printf("City: %s, Zip: %d\n", scan_city, scan_zip)

	var scan_student string
	var scan_gpa float64

	fmt.Print("Enter name and GPA: ")
	fmt.Scanf("%s %f", &scan_student, &scan_gpa)
	fmt.Printf("Name: %s, GPA: %.1f\n", scan_student, scan_gpa)

	// Error formatting functions
	err := checkAge(15)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive", age)
	}
	return nil
}
