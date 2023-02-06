package main

import (
	"bufio"
	"io/ioutil"

	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type AllItineraries struct {
	Itineraries map[string]Itinerarie `json:"Itineraries"`
}

type Itinerarie struct {
	Location    string `json:"Location"`
	DurOfTravel int    `json:"Duration"`
	StartDate   string `json:"Start Date"`
	EndDate     string `json:"End Date"`
}

type WeatherForecast struct {
	Name string `json:"Name"`
	Main struct {
		Temp float64 `json:"Temp"`
	} `json:"Main"`
}

// Starting main console
func main() {
outer:
	for {
		fmt.Println("                                                ")
		fmt.Println("[Welcome to the Travel Planner Console]\n\n",
			"(1) Proceed to itinerarie\n",
			"(2) Weather Update\n",
			"(3) Airbnb\n",
			"(4) Quit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1: //Create Itinerarie
			itinerariemain()
		case 2: //Weather Update
			weatherForecast()
		case 3: //Airbnb
		case 4: //quit

			break outer
		}
	}
}

// Itinerarie main page
func itinerariemain() {
outer:
	for {
		fmt.Println("                                                ")
		fmt.Println("[You are currently in Itinerarie Management Console]\n",
			"(1) View of itinerarie\n",
			"(2) Create of itinerarie\n",
			"(3) Update itinerarie (Duration, Start & End Date)\n",
			"(4) Back")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1: //View all itinerarie
			itinerarielist()
		case 2: //Creating of itinerarie
			itinerariecreate()
		case 3: // edit & updating itinerarie
			itinerariesupdate()
		case 4: //Return to main page
			break outer
		}
	}
}

// Itinerarie list display
func itinerarielist() {

	//Conneting to MYSQL Database 'db_itinerarie'
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db_itinerarie")

	//Handle error
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Retrieving itinerarie information from database
	results, err := db.Query("SELECT Location, Duration, StartDate, EndDate FROM itinerarie")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("                          ")
	fmt.Println("Itinerarie being retrieved")
	fmt.Println("===================================================")

	for results.Next() {
		var itinerarie Itinerarie

		err = results.Scan(&itinerarie.Location, &itinerarie.DurOfTravel, &itinerarie.StartDate, &itinerarie.EndDate)
		if err != nil {
			panic(err.Error())
		}

		//Display database retrieved

		fmt.Println("                                     ")
		fmt.Println(" Location: "+itinerarie.Location+"\n",
			"Duration of days:", itinerarie.DurOfTravel, "\n",
			"Start Date:", itinerarie.StartDate, "\n",
			"End Date: "+itinerarie.EndDate)

	}
	fmt.Println("===================================================")

}

// Creating of itinerarie
func itinerariecreate() {
	var newItineraries Itinerarie

	fmt.Print("\nPlease enter location (No space is required): ")
	fmt.Scanf("%v\n", &newItineraries.Location)

	fmt.Print("Please enter the duration of travel (Days): ")
	fmt.Scanf("%d\n", &(newItineraries.DurOfTravel))

	fmt.Print("Please enter start date (dd/mm/yyyy): ")
	reader0 := bufio.NewReader(os.Stdin)
	input0, _ := reader0.ReadString('\n')
	newItineraries.StartDate = strings.TrimSpace(input0)

	fmt.Print("Please enter end date (dd/mm/yyyy): ")
	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	newItineraries.EndDate = strings.TrimSpace(input1)

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db_itinerarie")

	//Handle error
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	}

	//Using SQL func'INSERT' to itinerarie into database
	insert, err := db.Query("INSERT INTO `db_itinerarie`.`Itinerarie` (`Location`, `Duration`, `StartDate`, `EndDate`) VALUES (?, ?, ?, ?)", newItineraries.Location, newItineraries.DurOfTravel, newItineraries.StartDate, newItineraries.EndDate)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("You have successfully added into the itineraries database")
}

// Updating of of itinerarie
func itinerariesupdate() {
	var editItineraries Itinerarie

	fmt.Print("\nPlease enter your location to be updated (No space is required): ")
	fmt.Scanf("%s\n", &editItineraries.Location)

	fmt.Print("Please enter the duration of travel: ")
	fmt.Scanf("%d\n", &editItineraries.DurOfTravel)

	fmt.Print("Please enter start date (dd/mm/yyyy): ")
	reader0 := bufio.NewReader(os.Stdin)
	input0, _ := reader0.ReadString('\n')
	editItineraries.StartDate = strings.TrimSpace(input0)

	fmt.Print("Please enter end date (dd/mm/yyyy): ")
	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	editItineraries.EndDate = strings.TrimSpace(input1)

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db_itinerarie")

	//Handle error
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	}

	result1 := db.QueryRow("select * from Itinerarie where Location= ? ", editItineraries.Location)

	err1 := result1.Scan(editItineraries.DurOfTravel, editItineraries.StartDate, editItineraries.EndDate)

	if err1 == sql.ErrNoRows {
		fmt.Println("This Location does not exist, please try again")
	} else {
		// SQL update function
		result, err := db.Exec("Update Itinerarie set Duration = ?, StartDate = ?, EndDate = ? where Location = ? ",
			editItineraries.DurOfTravel, editItineraries.StartDate, editItineraries.EndDate, editItineraries.Location)
		if err != nil {
			panic(err.Error())
		}
		id, err := result.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Itinerarie successful updated : %d \n", id)
		defer db.Close()
	}
}

// Display temperature of location entered
func weatherForecast() {
	var inputLocation string

	fmt.Print("\nPlease enter location: ")
	fmt.Scanf("%v\n", &inputLocation)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5010/weather/"+inputLocation, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				var res WeatherForecast
				json.Unmarshal(body, &res)
				if res.Name == "" {
					fmt.Println("Country not found!")
				} else {
					fmt.Println("Temperature in", res.Name, ":", res.Main.Temp, "°C")
				}
			}
		}
	}
}
