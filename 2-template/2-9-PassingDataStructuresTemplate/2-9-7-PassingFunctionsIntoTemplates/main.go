package main

import(
	"log"
	"os"
	"strings"
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


type AllData struct{
	NameInfo []Director
	OtherInfo []Info
}

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var tpl *template.Template
func init(){
	// New allocates a new, undefined template with the given name.
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string{
	// TrimSpace returns a slice of the string s, with all leading
	// and trailing white space removed, as defined by Unicode.
	// Fast path for ASCII: look for the first ASCII non-space byte
	s = strings.TrimSpace(s)
	if len(s)>3{
		s = s[:3]
	}
	return s
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

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err!=nil{
		log.Fatalln(err)
	}
}
