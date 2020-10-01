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
	data := []string{"Kiarostami", "Polansky", "Haneke", "Wong Kar Wei"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err!=nil{
		log.Fatalln(err)
	}
}
