package main

import(
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main(){
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Hey, we are checking for the variables in template")
	if err!=nil{
		log.Fatalln(err)
	}
}