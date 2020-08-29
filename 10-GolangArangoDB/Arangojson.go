package main

import(
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	//"strconv"
)

type Flight struct{
	_key string `json:"key"`
	_id string `json:"id"`
	_from string `json:"from"`
	_to string `json:"to"`
	Rev string `json:"rev"`
	Year int32 `json:"year"`
	Month int32 `json:"month"`
	Day int32 `json:"day"`
	DayOfWeek int32 `json:"dayofweek"`
	DepTime int32 `json:"deptime"`
	ArrTime int32 `json:"arrtime"`
	DepTimeUTC string `json:"deptimeutc"`
	ArrTimeUTC string `json:"arrtimeutc"`
	UniqueCarrier string `json:"uniquecarrier"`
	FlightNum int32 `json:"flightnum"`
	TailNum string `json:"tailnum"`
	Distance int32 `json:"distance"`
}

type Flights struct {
	Flights []Flight `json:"flights"`
}


func main(){
	// Open our jsonFile
	jsonFile, err := os.Open("results-_system.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	//Parsing with Structs
	//We have a few options when it comes to parsing the JSON that is contained within our users.json file. We could
	//either unmarshal the JSON using a set of predefined structs, or we could unmarshal the JSON using a map[string]
	//interface{} to parse our JSON into strings mapped against arbitrary data types.

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Flights array
	var flights []Flight

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'flights' which we defined above
	err = json.Unmarshal(byteValue, &flights)
	if err != nil{
		fmt.Println(err)
	}


	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < 5; i++ {
		fmt.Println(flights[i])
		fmt.Println("ID: " + flights[i]._id)
		//fmt.Println("User Age: " + strconv.Itoa(flights.Flights[i].Day))
		fmt.Println("From: " + flights[i]._from)
		fmt.Println("To: " + flights[i]._to)
	}
}
