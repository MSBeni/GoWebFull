package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	//"reflect"
)

type TstUser struct {
	Username string
	Identifier string
	Onetime string
	password string
	Recovery string
	code string
	First string
}

type User struct{
	FirstName string
	LastName string
	CompanyName string
	Address string
	City string
	Country string
	State string
	Zip string
	Phone1 string
	Phone2 string
	Email string
	Web string
}


func main() {
	// Open the file
	csvfile, err := os.Open("Test.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close() // this needs to be after the err check


	lines, err := csv.NewReader(csvfile).ReadAll()

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}


	fmt.Println(lines)
	var usersall []TstUser
	//var usersall []User

	for idx, line := range lines {
		if idx == 0 {
			continue
		}
		//tstuser := User{
		//	FirstName:       line[1],
		//	LastName:        line[2],
		//	CompanyName:     line[3],
		//	Address:         line[4],
		//	City:            line[5],
		//	Country:         line[6],
		//	State:           line[7],
		//	Zip:             line[8],
		//	Phone1:          line[9],
		//	Phone2:          line[10],
		//	Email:           line[11],
		//	Web:             line[12],
		//}
		tstuser := TstUser{
			Username:       line[1],
			Identifier:     line[2],
			Onetime:        line[3],
			password:       line[4],
			Recovery:       line[5],
			code:           line[6],
			First:          line[7],
		}
		usersall = append(usersall, tstuser)
	}

	fmt.Println(usersall)
	//reflect.TypeOf(x).Kind()

	//for {
	//	// Read each record from csv
	//	record, err := csv.NewReader(csvfile).Read()
	//	if err == io.EOF {
	//		break
	//	}
	//
	//	if err != nil {
	//		log.Println(err)
	//		log.Println("Ignore")
	//	} else {
	//		fmt.Println(record)
	//	}
	//}
}
