package main

import (
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

func encode(w http.ResponseWriter, r *http.Request) {

}
func decode(w http.ResponseWriter, r *http.Request) {

}
