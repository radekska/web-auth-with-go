package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name string
}

func main() {
	p1 := person{
		Name: "Jenny",
	}

	p2 := person{
		Name: "Joe",
	}

	personList := []person{p1, p2}
	marshalled, err := json.Marshal(personList)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(marshalled))

}
