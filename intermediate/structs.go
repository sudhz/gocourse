package intermediate

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

type Address struct {
	city    string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

func main() {
	// Normal structs
	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "London",
			country: "UK",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "3420938543",
			cell: "2034934284",
		},
	}

	p2 := Person{
		firstName: "Jane",
		age:       25,
	}
	p2.address.city = "New York"
	p2.address.country = "USA"

	p3 := Person{
		firstName: "Jane",
		age:       25,
		address: Address{
			city:    "New York",
			country: "USA",
		},
	}

	fmt.Printf("Before increment, p1:  %+v\n", p1)
	fmt.Printf("p2: %+v\n", p2)
	fmt.Printf("p1 fullName: %s\n", p1.fullName())

	p1.incrementAgeByOne()
	fmt.Printf("After increment, p1:  %+v\n", p1)

	fmt.Printf("p1.home: %s\n", p1.home)
	fmt.Printf("p1.cell: %s\n", p1.cell)
	fmt.Printf("p1.PhoneHomeCell.home: %s\n", p1.PhoneHomeCell.home)
	fmt.Printf("p1.PhoneHomeCell.cell: %s\n", p1.PhoneHomeCell.cell)

	fmt.Println("Are p2 and p3 equal?:", p2 == p3)

	// Anonymous structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoemail@example.org",
	}

	fmt.Printf("user: %+v\n", user)
}
