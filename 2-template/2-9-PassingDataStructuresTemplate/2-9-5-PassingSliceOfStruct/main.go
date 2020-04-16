package main

import(
	"log"
	"os"
	"text/template"
)

type Director struct {
	Name string
	Known string
}
var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseGlob("tpl2.gohtml"))
}
func main(){
	data := Director{
		Name:  "Kiarostami",
		Known: "Reality",
	}
	//dataa := map[string]string{"Kiarostami":"Reality", "Polansky":"Gossip", "Haneke":"Psycho", "Wong Kar Wei":"Love"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", data)
	if err!=nil{
		log.Fatalln(err)
	}
}
