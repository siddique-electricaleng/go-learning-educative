package person

type Person struct {
	firstName string
	lastName  string
}

// Getter - simply return the firstName and lastName separately

func (p *Person) GetFirstName() string {
	return p.firstName
}
func (p *Person) GetLastName() string {
	return p.lastName
}

// Setter method - set the firstName and LastName separately

func (p *Person) SetFirstName(name string) {
	p.firstName = name
}
func (p *Person) SetLastName(name string) {
	p.lastName = name
}
