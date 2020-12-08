package phonebook

type Person struct {
	name    string
	age     int16
	gender  bool
	address string
}

// PhoneBook definition
type PhoneBook struct {
	persons []Person
}

func NewPhoneBook() *PhoneBook {
	pb := &PhoneBook{}
	return pb
}

func (pb *PhoneBook) AddPerson() *PhoneBook {
	
}