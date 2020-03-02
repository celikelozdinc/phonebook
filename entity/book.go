package entity

import "fmt"

//PhoneBook is used to store PhoneRecord
type PhoneBook struct {
	RecordList []*PhoneRecord
}

//Push apends new records to phonebook
func (book *PhoneBook) Push(record *PhoneRecord) {
	book.RecordList = append(book.RecordList, record)
}

//Printer will print contents
func (book *PhoneBook) Printer() {
	fmt.Println("... Records in phonebook ...")
	for index, _ := range book.RecordList {
		book.RecordList[index].Printer()
	}
	fmt.Println("... Records in phonebook ...")
}

//SearchByName searches using name of the record
func (book *PhoneBook) SearchByName(name string) *PhoneRecord {

	i := -1

	for index, element := range book.RecordList {
		if element.Name == name {
			i = index
		}
	}

	if i != -1 {
		fmt.Println("Record could be found.")
		return book.RecordList[i]
	}

	// if does not found, returns an empty record
	fmt.Println("Record could not be found.")
	return &PhoneRecord{}

}
