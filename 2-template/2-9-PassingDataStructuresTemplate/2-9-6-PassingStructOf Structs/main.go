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
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}
func main(){
	Abbas := Director{
		Name:  "Kiarostami",
		Known: "Reality",
	}
	Roman := Director{
		Name:  "Polansky",
		Known: "Gossip",
	}
	Michele := Director{
		Name:  "Haneke",
		Known: "Psycho",
	}
	Wong := Director{
		Name:  "Wong Kar Wei",
		Known: "Love",
	}
	data := []Director{Abbas, Roman, Michele, Wong}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	//err := tpl.Execute(os.Stdout, data)
	if err!=nil{
		log.Fatalln(err)
	}
}
