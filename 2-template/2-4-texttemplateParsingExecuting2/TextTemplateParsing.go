package main

import (
	"log"
	"os"
	"text/template"
)

func main(){
	tpl,err := template.ParseFiles("tpl.gohtml")
	if err!=nil{
		log.Fatalln(err)
	}
	fn, err := os.Create("tpl.gohtml")
	if err!=nil{
		log.Fatalln(err)
	}
	defer fn.Close()

	err = tpl.Execute(fn, nil)
	if err!=nil{
		log.Fatalln(err)
	}
}
