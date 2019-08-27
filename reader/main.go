package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

//Method in charge to read and parse the file and also invoke the other operations
func read(fileName string, action ActionDB, post CrmPost) (int, error) {
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	i := 0
	for {
		line, error := reader.Read()
		if error == io.EOF {
			return i, nil
		} else if error != nil {
			return -1, errors.New(error.Error())
		}
		customer := Customer{
			Id:        line[0],
			Firstname: line[1],
			Lastname:  line[2],
			Email:     line[3],
			Phone:     line[4],
		}
		//insert statement against the DB
		action(customer, insert)
		customerJson, _ := json.Marshal(customer)
		//HTTP request against the container which holds the CRM
		post(customerJson)
		i++
	}
}

func main() {
	//Root url runs the CSV parser
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		_, err := read("/tmp/reader/MOCK_DATA.csv", action, post)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		fmt.Fprintf(w, "CSV Parsed")
	})

	http.ListenAndServe(":5000", nil)
}
