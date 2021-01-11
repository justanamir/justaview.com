package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// data := struct {
	// 	Name string
	// 	Age  int
	// }{"Amir Danial", 30}

	// data := map[string]string{
	// 	"Name": "Amir Danial",
	// 	"Age":  "30",}

	type contact struct {
		Email     string
		Telephone string
	}

	type details struct {
		Name string
		Age  int
		contact
	}

	data := details{
		Name: "Amir Danial",
		Age:  30,
		contact: contact{
			Email:     "amirdanialmuslim@gmail.com",
			Telephone: "018-2547003",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
