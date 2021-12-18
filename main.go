package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name string
}

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		print(err)
	}
}

func encode(writer http.ResponseWriter, request *http.Request) {
	p1 := person{
		Name: "Jenny",
	}
	err := json.NewEncoder(writer).Encode(p1)
	if err != nil {
		exceptionMessage("encode", err)
	}
}
func decode(writer http.ResponseWriter, request *http.Request) {
	var p1 person
	err := json.NewDecoder(request.Body).Decode(&p1)
	if err != nil {
		exceptionMessage("decode", err)
	} else {
		log.Println("Person:", p1)
	}
}

func exceptionMessage(message string, err error) {
	log.Printf("Tried to %s invalid data.\n", message)
	log.Println(err)
}
