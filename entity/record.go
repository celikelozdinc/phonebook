package entity

import "fmt"

// PhoneRecord holds records for PhoneBook
type PhoneRecord struct {
	Phone   string
	Name    string
	Surname string
	Country string
	ID      int64
}

//Printer will print contents
//type *PhoneRecord is the receiver of the Printer()
func (ptr *PhoneRecord) Printer() {
	fmt.Printf("%d: %s %s - %s - %s\n", ptr.ID, ptr.Name, ptr.Surname, ptr.Country, ptr.Phone)
}


func (ptr *PhoneRecord) UpdateID() {
	ptr.ID = 0
}