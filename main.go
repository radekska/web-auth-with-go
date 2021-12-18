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
	//p1 := person{
	//	Name: "Jenny",
	//}
	//
	//p2 := person{
	//	Name: "Joe",
	//}
	//
	//personList := []person{p1, p2}
	//marshalled, err := json.Marshal(personList)
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Println(string(marshalled))
	//
	//var personList2 []person
	//err = json.Unmarshal(marshalled, &personList2)
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Println(personList2)

	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	http.ListenAndServe(":8080", nil)

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
