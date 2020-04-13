package main

import(
	"fmt"
)

type person struct {
	Fname string
	Name string
	Age int32
}

type Authority struct{
	person
	AbleToEnterTheClub bool
}

type HumanInPower interface {
	speak()
	listen()
}

func (p person) speak(){
	fmt.Println(p.Name, p.Fname,`Saying Hello, as simple as it can be`)
}

func (Ap Authority) speak(){
	fmt.Println(Ap.Name, Ap.Fname,`has the authority to speak`)
}

func (AP Authority) listen(){
	agent := AP.Fname + AP.Fname
	fmt.Println("agent in power of listening", agent, "in age:", AP.Age)
}

func ShowAbilities(h HumanInPower){
	h.speak()
	h.listen()
}

func main(){
	a := person{
		Fname: "Micky",
		Name:  "John",
		Age:   23,
	}

	aa := Authority{
		person:             person{"SBeni", "Mo", 29},
		AbleToEnterTheClub: true,
	}
	a.speak()
	aa.speak()
	ShowAbilities(aa)

}
