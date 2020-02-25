package entity


import "fmt"


type PhoneBook struct {
	RecordList []*PhoneRecord
}

func (book *PhoneBook) Push(record *PhoneRecord){
	book.RecordList = append(book.RecordList, record)
}

func (book *PhoneBook) Printer(){
	fmt.Println("... Records in phonebook ...")
	for index,_ := range book.RecordList{
		book.RecordList[index].Printer()
	}
	fmt.Println("... Records in phonebook ...")
}


func (book *PhoneBook) SearchByName(name string) *PhoneRecord{

	i := -1

	for index, element := range book.RecordList {
		if element.Name == name {
			i = index
		}
	}

	return book.RecordList[i]

}
