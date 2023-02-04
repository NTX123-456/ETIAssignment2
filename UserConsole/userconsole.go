package main

import (
	"bufio"
	"bytes"

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

// Starting main console
func main() {
outer:
	for {
		fmt.Println("                                                ")
		fmt.Println("[Welcome to the Travel Planner Console]\n\n",
			"(1) Proceed to itinerarie\n",
			"(2) For Weather updates\n",
			"(3) Airbnb\n",
			"(4) Quit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1: //Create Itinerarie
			itinerariemain()
		case 2: //Weather Update
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
			"(3) Edit & Update itinerarie\n",
			"(4) Back")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1: //List all itinerarie
			itinerarielist()
		case 2: //Creating itinerarie
			itinerariecreate()
		case 3: // edit & updating itinerarie
		case 4: //Return to main page
			break outer
		}
	}
}

// Itinerarie list display
func itinerarielist() {

	//Conneting to MYSQL Database 'my_db'
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db_itinerarie")

	//Handle error
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Retrieving driver's information from database
	results, err := db.Query("SELECT Location, Duration, StartDate, EndDate FROM itinerarie")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("                          ")
	fmt.Println("Itinerarie being retrieved")
	fmt.Println("___________________________")

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
}

// Itinerarie creating an account
func itinerariecreate() {
	var newItineraries Itinerarie
	//var Location string

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

	//Using SQL func'INSERT' to add account into database
	insert, err := db.Query("INSERT INTO `db_itinerarie`.`Itinerarie` (`Location`, `Duration`, `StartDate`, `EndDate`) VALUES (?, ?, ?, ?)", newItineraries.Location, newItineraries.DurOfTravel, newItineraries.StartDate, newItineraries.EndDate)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("You have successfully added into the itineraries database")
}

// Itinerarie updating account
func passengerupdate() {
	var newItineraries Itinerarie
	var Location string

	fmt.Print("\nPlease enter location: ")
	fmt.Scanf("%v\n", &Location)

	fmt.Print("Please enter the duration of travel: ")
	fmt.Scanf("%d\n", &(newItineraries.DurOfTravel))

	fmt.Print("Please enter start date: ")
	reader0 := bufio.NewReader(os.Stdin)
	input0, _ := reader0.ReadString('\n')
	newItineraries.StartDate = strings.TrimSpace(input0)

	fmt.Print("Please enter end date: ")
	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	newItineraries.EndDate = strings.TrimSpace(input1)

	jsonString, _ := json.Marshal(newItineraries)
	resbody := bytes.NewBuffer(jsonString)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/itineraries/"+Location, resbody); err == nil {
		if res, err := client.Do(req); err == nil {

			if res.StatusCode == 202 {
				fmt.Println("Itinerarie", Location, "is being updated")
			} else if res.StatusCode == 409 {
				fmt.Println("Error - itinerarie", Location, "is not being updated")
			}
		}
	}
}
