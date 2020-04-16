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
	data := map[string]string{"Kiarostami":"Reality", "Polansky":"Gossip", "Haneke":"Psycho", "Wong Kar Wei":"Love"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err!=nil{
		log.Fatalln(err)
	}
}
