package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"strings"
)

func main() {
	const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}

	log.Println("Data Awal => ", reflect.TypeOf(jsonStream))
	log.Println("Reader => ", strings.NewReader(jsonStream))
	log.Println("Type Reader => ", reflect.TypeOf(strings.NewReader(jsonStream)))
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	log.Println("*json.Decoder FILE => ", dec)
	log.Println("Type Data => ", reflect.TypeOf(dec))
	for {
		var m Message
		log.Println("MESSAGE =>", m)
		if err := dec.Decode(&m); err == io.EOF {
			log.Println("ERRROOOR CUK", reflect.TypeOf(err))
			log.Println("===========", io.EOF, reflect.TypeOf(io.EOF))
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Result => %s: %s\n", m.Name, m.Text)
	}
}
