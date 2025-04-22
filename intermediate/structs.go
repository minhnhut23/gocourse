package intermediate

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneCell
}

type Address struct {
	city    string
	country string
}

type PhoneCell struct {
	home string
	cell string
}

func main() {

	p1 := Person{
		firstName: "a",
		lastName:  "b",
		age:       12,
		address: Address{
			city:    "abc",
			country: "def",
		},
		PhoneCell: PhoneCell{
			home: "sdfads",
			cell: "054829584",
		},
	}
	p2 := Person{
		lastName: "c",
		age:      22,
	}

	p2.address.city = "adfs"
	p2.address.country = "ádffdsa"

	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p1.fullName())
	fmt.Println(p1.address)
	fmt.Println(p2.address)
	fmt.Println(p1.PhoneCell)

	user := struct {
		userName string
		email    string
	}{
		userName: "user123",
		email:    "a@mail.com",
	}

	fmt.Println(user)

	fmt.Println("Before increment:", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After increment:", p1.age)
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}
