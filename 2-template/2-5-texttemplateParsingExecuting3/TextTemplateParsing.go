package main

import(
	"os"
	"text/template"
	"log"
)

func main(){
	tpl, err := template.ParseFiles("one.gohtml")
	if err!=nil{
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	tpl, err = tpl.ParseFiles("vespa.gohtml", "one.gohtml", "two.gohtml")
	if err!=nil{
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err!=nil{
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err!=nil{
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err!=nil{
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err!=nil{
		log.Fatalln(err)
	}



}