package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main(){
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	tpl := fmt.Sprint(	`<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!!!</title>
	</head>
	<body>
	<h1>`+name+`</h1>
	</html>
		`)

	nf, err := os.Create("index.html")
	if err!=nil{
		log.Fatal("error creating the index.html file")
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))

}