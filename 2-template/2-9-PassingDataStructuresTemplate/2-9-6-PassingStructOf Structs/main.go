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

type Info struct {
	Age int
	Alive bool
}

type InfoTea struct {
	Brand string
	Base string
}

type AllData struct{
	NameInfo []Director
	OtherInfo []Info
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

	AbbasKia := Info{
		Age:  65,
		Alive: false,
	}
	RomanPo := Info{
		Age:  60,
		Alive: true,
	}
	MicheleHa := Info{
		Age:  82,
		Alive: true,
	}
	WongWei := Info{
		Age:  55,
		Alive: true,
	}
	Dire := []Director{Abbas, Roman, Michele, Wong}
	OthInf := []Info{AbbasKia, RomanPo, MicheleHa, WongWei}
	data := AllData{
		NameInfo:  Dire,
		OtherInfo: OthInf,
	}

	err := tpl.Execute(os.Stdout, data)
	if err!=nil{
		log.Fatalln(err)
	}
}
