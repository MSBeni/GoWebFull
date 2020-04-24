package main

import (
	"GoWebFull/3-Building_Simple_API_Web_Server/handlers"
	"log"
	"net/http"
)


func main(){
	// Checking for the handler function with http.HandleFunc command
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/", handlers.RootHandler)
	// running on the local host port defined
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil{
		log.Fatalln(err)
	}
}