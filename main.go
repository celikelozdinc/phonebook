package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	entities "github.com/celikelozdinc/phonebook/entity"
)

func search(book *entities.PhoneBook, searchStr string, c chan *entities.PhoneRecord, wg *sync.WaitGroup) {
	defer wg.Done()
	record := book.SearchByName(searchStr)
	c <- record
}

func prepare(id int64, line []string, c chan *entities.PhoneRecord, wg *sync.WaitGroup) {
	defer wg.Done()
	newRecord := &entities.PhoneRecord{
		ID:      id,
		Name:    line[1],
		Surname: line[2],
		Country: line[3],
		Phone:   line[4],
	}

	c <- newRecord
}

func main() {

	fmt.Println("... Starting Main package...")

	csvFile, _ := os.Open("data/records.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	//Initialize slice
	//var Records []*entity.PhoneRecord
	Records := []*entities.PhoneRecord{}

	//Initialize waitgroup
	//how many goroutines need to be waited
	var wg sync.WaitGroup
	wg.Add(3)

	//use channel in order to create record OR search the record
	c := make(chan *entities.PhoneRecord)

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		id, _ := strconv.ParseInt(line[0], 10, 64)

		//use channel in order to append a record
		go prepare(id, line, c, &wg)
		newRecord := <-c
		Records = append(Records, newRecord)

	}

	MyBook := &entities.PhoneBook{
		RecordList: Records,
	}

	MyBook.Push(&entities.PhoneRecord{ID: 100, Name: "X", Surname: "Y", Country: "Z", Phone: "T"})
	MyBook.Push(&entities.PhoneRecord{ID: 200, Name: "A", Surname: "B", Country: "C", Phone: "D"})

	//use channel in order to search the record
	fmt.Println("... Record will be searched on phonebook ... ")
	go search(MyBook, "Simge", c, &wg)
	record := <-c

	close(c)

	record.Printer()

	MyBook.Printer()

	wg.Wait()
	fmt.Println("... Finishing Main package...")
}
