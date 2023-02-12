package main

import (
	"bufio"
	"log"
	"net/http"

	"database/sql"
	"fmt"
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

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Welcome to the Hotel Console\n</h1>")

}

// Itinerarie main console page
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
		case 1: // View all itineraries
			itinerarielist()
		case 2: // Creating of itineraries
			itinerariecreate()
		case 3: // edit & updating itinerarie
			itinerariesupdate()
		case 4: // Return to main page
			break outer
		}
	}
}

// Display all Itineraries through database (db_itinerarie)
func itinerarielist() {

	// Conneting with MYSQL Database 'db_itinerarie'
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db_itinerarie")

	// Handle error
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Using SQL command 'SELECT' to retrieve itineraries from database
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

		// Display database information being retrieved
		fmt.Println("                                     ")
		fmt.Println(" Location: "+itinerarie.Location+"\n",
			"Duration of days:", itinerarie.DurOfTravel, "\n",
			"Start Date:", itinerarie.StartDate, "\n",
			"End Date: "+itinerarie.EndDate)

	}
	fmt.Println("===================================================")

}

// Creating of new itinerarie
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

	// Handle error
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	}

	// Using SQL command 'INSERT' to create new itinerarie into database
	insert, err := db.Query("INSERT INTO `db_itinerarie`.`Itinerarie` (`Location`, `Duration`, `StartDate`, `EndDate`) VALUES (?, ?, ?, ?)", newItineraries.Location, newItineraries.DurOfTravel, newItineraries.StartDate, newItineraries.EndDate)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("You have successfully added into the itineraries database")
}

// Update existing itinerarie
func itinerariesupdate() {
	var editItineraries Itinerarie

	fmt.Print("\nPlease enter your location to be updated (No space is required): ")
	fmt.Scanf("%s\n", &editItineraries.Location)

	fmt.Print("Please enter the duration of travel to be updated: ")
	fmt.Scanf("%d\n", &editItineraries.DurOfTravel)

	fmt.Print("Please enter start date to be updated(dd/mm/yyyy): ")
	reader0 := bufio.NewReader(os.Stdin)
	input0, _ := reader0.ReadString('\n')
	editItineraries.StartDate = strings.TrimSpace(input0)

	fmt.Print("Please enter end date to be updated(dd/mm/yyyy): ")
	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	editItineraries.EndDate = strings.TrimSpace(input1)

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db_itinerarie")

	// Handle error
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
		// Using SQL command 'UPDATE' to update existing itinerarie within the database
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
