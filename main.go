package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"log"
	"encoding/csv"
	entities "phonebook/entity"
	"strconv"
)

func main() {

	fmt.Println("..Main package...")

	csvFile, _ := os.Open("data/records.csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))


	//Initialize slice
	//var Records []*entity.PhoneRecord
	Records := []*entities.PhoneRecord{}


	for 
	{
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
		}
		id, _:= strconv.ParseInt(line[0],10,64)

		Records = append(Records, &entities.PhoneRecord{
            ID: id,
			Name:  line[1],
			Surname: line[2],
			Country: line[3],
			Phone: line[4],
            },
		)
	}


	MyBook := &entities.PhoneBook{
		RecordList: Records,
	}

	MyBook.Push(&entities.PhoneRecord{ID:100,Name:"X",Surname:"Y",Country:"Z",Phone:"T"})
	MyBook.Push(&entities.PhoneRecord{ID:200,Name:"A",Surname:"B",Country:"C",Phone:"D"})

	record := MyBook.SearchByName("X")
	record.Printer()


	MyBook.Printer()

}

