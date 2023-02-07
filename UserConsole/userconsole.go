package main

import (
	"bufio"
	"errors"
	"io/ioutil"
	"strconv"

	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"gorm.io/gorm"

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

type Hotel struct {
	ID             int    `json:ID`
	HotelName      string `json:Hotel Name"`
	HotelInfo      string `json:"Hotel Information"`
	HotelAddr      string `json:"Hotel Address"`
	HotelStar      int    `json:"Hotel Stars"`
	HotelAmenities string `json:"Hotel Amenities"`
	Price          int    `json:"Hotel Price"`
	Country        string `json:"Hotel Amenities"`
}

type PackageModel struct {
	gorm.Model
	Name        string `json:"name"`
	PackageType string `json:"package_type"`
	Price       int    `json:"price"`
}

func (p PackageModel) TableName() string {
	return "packages"
}

type BookingModel struct {
	gorm.Model
	PackageId int    `json:"package_id"`
	Username  string `json:"username"`
}

func (p BookingModel) TableName() string {
	return "bookings"
}

type BookedPackage struct {
	BookingId       int    `json:"booking_id"`
	BookingUsername string `json:"booking_username"`
	PackageId       int    `json:"package_id"`
	PackageName     string `json:"package_name"`
	PackageType     string `json:"package_type"`
	PackagePrice    int    `json:"package_price"`
}

/** End: Models **/

func clear() {
	fmt.Print("\033[2J") // clear
	fmt.Print("\033[H")  // reset cursor
}

func printTitle() {
	clear()
	fmt.Println("---------------------")
	fmt.Println(" Booking App")
	fmt.Println("---------------------")
	fmt.Println(" Type 'main' to return to main menu.")
	fmt.Println(" Type 'quit' to exit the program.")
	fmt.Println("---------------------")
}

func printSubtitle(subtitle string) {
	fmt.Printf(" %s\n", subtitle)
	fmt.Println("---------------------")
}

func printStatusMessage(message string) {
	fmt.Printf(" > %s\n", message)
	fmt.Println("---------------------")
}

func printErrStatus(err error) {
	printStatusMessage(fmt.Sprintf("please try again, an error occurred: '%v'", err))
}

func printBookingMenu() {
	printTitle()
	fmt.Println(" What you want to do?")
	fmt.Println(" 1. List all tourist attractions and hotels")
	fmt.Println(" 2. Book tourist attraction or hotel")
	fmt.Println(" 3. Cancel or update booking")
}

func getInput() string {
	fmt.Print(" -> ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(strings.Replace(text, "\n", "", -1))
	return text
}

func enterToContinue() {
	fmt.Println(" Press 'Enter' to return to main menu")
	_, _ = reader.ReadString('\n')
}

var reader = bufio.NewReader(os.Stdin)

var username string

// Starting main console
func main() {
outer:
	for {
		fmt.Println("                                                ")
		fmt.Println("[Welcome to the Travel Planner Console]\n\n",
			"(1) Planning of itinerarie\n",
			"(2) Weather Update\n",
			"(3) View Hotels\n",
			"(4) Booking Hotel & Attractions\n",
			"(5) Quit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1: // Itinerarie features
			itinerariemain()
		case 2: //Weather Update
			weatherForecast()
		case 3: //View Hotels
			HotelConsole()
		case 4: //Booking Hotel/Attractions
			BookingMenu()
		case 5: //Quit
			break outer
		}
	}
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

func HotelConsole() {
outer:
	for {
		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("You are Now In the Hotel Console")
		fmt.Println("Get Hotels\n",
			"1. List all Hotels\n",
			"2. List Hotels by Country\n",
			"3. List Hotels by Hotel rating\n",
			"4. List Hotels by Hotel Amenities\n",
			"5. List Hotels by Price\n",
			"6. Quit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			getallHotels()
		case 2:
			getHotelsbyCountry()
		case 3:
			getHotelsbyStar()
		case 4:
			getHotelsbyAmenities()
		case 5:
			getHotelsByPrice()
		case 6:
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

func getHotelsbyCountry() {
	var h Hotel

	fmt.Println("Enter which country to display hotels from: ")
	fmt.Scanf("%s\n", &h.Country)

	db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/ETIASSG2_db")
	if err != nil {
		panic(err.Error())

	}
	result1 := db.QueryRow("select * from Hotels where Country = ?", &h.Country)
	err1 := result1.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
	if err1 == sql.ErrNoRows {
		fmt.Println("Sorry we do not have hotels from that country")
		BookingMenu()
	} else {

		results, err := db.Query("select * from Hotels where Country = ?", h.Country)
		if err != nil {
			fmt.Println("Error")
			panic(err.Error())

		}
		for results.Next() {

			err = results.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("\nHotel Name: "+h.HotelName, "\nHotel Information: "+h.HotelInfo, "\nHotel Address: "+h.HotelAddr, "\nHotel Star: ", h.HotelStar, "\nHotel Amenities: "+h.HotelAmenities, "\nHotel Price: ", h.Price, "\nCountry: "+h.Country)
		}

		defer db.Close()

	}
}

func getHotelsbyStar() {
	var h Hotel

	fmt.Println("Enter the amount of rating of the hotel: ")
	fmt.Scanf("%d\n", &h.HotelStar)

	db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/ETIASSG2_db")
	if err != nil {
		panic(err.Error())

	}
	result1 := db.QueryRow("select * from Hotels where HotelStar = ?", &h.HotelStar)
	err1 := result1.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
	if err1 == sql.ErrNoRows {
		fmt.Println("Sorry we do not have hotels from this star")
		BookingMenu()
	} else {

		results, err := db.Query("select * from Hotels where HotelStar = ?", &h.HotelStar)
		if err != nil {
			fmt.Println("Error")
			panic(err.Error())

		}
		for results.Next() {

			err = results.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("\nHotel Name: "+h.HotelName, "\nHotel Information: "+h.HotelInfo, "\nHotel Address: "+h.HotelAddr, "\nHotel Star: ", h.HotelStar, "\nHotel Amenities: "+h.HotelAmenities, "\nHotel Price: ", h.Price, "\nCountry: "+h.Country)
		}

		defer db.Close()

	}
}

func getHotelsbyAmenities() {
	var h Hotel

	fmt.Println("Enter the amenities of the hotel: ")
	fmt.Scanf("%s\n", &h.HotelAmenities)

	db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/ETIASSG2_db")
	if err != nil {
		panic(err.Error())

	}
	result1 := db.QueryRow("select * from Hotels where HotelAmenities LIKE ?", "%"+h.HotelAmenities+"%")
	err1 := result1.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
	if err1 == sql.ErrNoRows {
		fmt.Println("Sorry we do not have this type of Amenities")
		BookingMenu()
	} else {

		results, err := db.Query("select * from Hotels where HotelAmenities LIKE ?", "%"+h.HotelAmenities+"%")
		if err != nil {
			fmt.Println("Error")
			panic(err.Error())

		}
		for results.Next() {

			err = results.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("\nHotel Name: "+h.HotelName, "\nHotel Information: "+h.HotelInfo, "\nHotel Address: "+h.HotelAddr, "\nHotel Star: ", h.HotelStar, "\nHotel Amenities: "+h.HotelAmenities, "\nHotel Price: ", h.Price, "\nCountry: "+h.Country)
		}

		defer db.Close()

	}
}

func getHotelsByPrice() {
outer:
	for {
		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("Select the Price to Filter By\n")
		fmt.Println(
			"1. Get Hotels below $50\n",
			"2. Get Hotels between $50 and $100\n",
			"3. Get Hotels between $100 and $200\n",
			"4. Get Hotels between $200 and $300\n",
			"5. Get Hotels between $300 and $400\n",
			"6. Return to Hotel Console")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			getHotelBelow50()
		case 2:
			getHotelBw50and100()
		case 3:
			getHotelBw100and200()
		case 4:
			getHotelBw200and300()
		case 5:
			getHotelBw300and400()
		case 6:
			break outer
		}
	}
}

func getHotelBelow50() {
	var h Hotel

	db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/ETIASSG2_db")
	if err != nil {
		panic(err.Error())

	}
	result1 := db.QueryRow("select * from Hotels where price < 50")
	err1 := result1.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
	if err1 == sql.ErrNoRows {
		fmt.Println("Sorry we do not have hotels from this price range")
		BookingMenu()
	} else {

		results, err := db.Query("select * from Hotels where price < 50")
		if err != nil {
			fmt.Println("Error")
			panic(err.Error())

		}
		for results.Next() {

			err = results.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("\nHotel Name: "+h.HotelName, "\nHotel Information: "+h.HotelInfo, "\nHotel Address: "+h.HotelAddr, "\nHotel Star: ", h.HotelStar, "\nHotel Amenities: "+h.HotelAmenities, "\nHotel Price: ", h.Price, "\nCountry: "+h.Country)
		}

		defer db.Close()

	}
}

func getHotelBw50and100() {
	var h Hotel

	db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/ETIASSG2_db")
	if err != nil {
		panic(err.Error())

	}
	result1 := db.QueryRow("select * from Hotels where price > 50 and price < 100")
	err1 := result1.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
	if err1 == sql.ErrNoRows {
		fmt.Println("Sorry we do not have hotels from this price range")
		BookingMenu()
	} else {

		results, err := db.Query("select * from Hotels where price > 50 and price < 100")
		if err != nil {
			fmt.Println("Error")
			panic(err.Error())

		}
		for results.Next() {

			err = results.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("\nHotel Name: "+h.HotelName, "\nHotel Information: "+h.HotelInfo, "\nHotel Address: "+h.HotelAddr, "\nHotel Star: ", h.HotelStar, "\nHotel Amenities: "+h.HotelAmenities, "\nHotel Price: ", h.Price, "\nCountry: "+h.Country)
		}

		defer db.Close()

	}
}

func getHotelBw100and200() {
	var h Hotel

	db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/ETIASSG2_db")
	if err != nil {
		panic(err.Error())

	}
	result1 := db.QueryRow("select * from Hotels where price > 100 and price < 200")
	err1 := result1.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
	if err1 == sql.ErrNoRows {
		fmt.Println("Sorry we do not have hotels from this price range")
		BookingMenu()
	} else {

		results, err := db.Query("select * from Hotels where price > 100 and price < 200")
		if err != nil {
			fmt.Println("Error")
			panic(err.Error())

		}
		for results.Next() {

			err = results.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("\nHotel Name: "+h.HotelName, "\nHotel Information: "+h.HotelInfo, "\nHotel Address: "+h.HotelAddr, "\nHotel Star: ", h.HotelStar, "\nHotel Amenities: "+h.HotelAmenities, "\nHotel Price: ", h.Price, "\nCountry: "+h.Country)
		}

		defer db.Close()

	}
}

func getHotelBw200and300() {
	var h Hotel

	db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/ETIASSG2_db")
	if err != nil {
		panic(err.Error())

	}
	result1 := db.QueryRow("select * from Hotels where price > 200 and price < 300")
	err1 := result1.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
	if err1 == sql.ErrNoRows {
		fmt.Println("Sorry we do not have hotels from this price range")
		BookingMenu()
	} else {

		results, err := db.Query("select * from Hotels where price > 200 and price < 300")
		if err != nil {
			fmt.Println("Error")
			panic(err.Error())

		}
		for results.Next() {

			err = results.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("\nHotel Name: "+h.HotelName, "\nHotel Information: "+h.HotelInfo, "\nHotel Address: "+h.HotelAddr, "\nHotel Star: ", h.HotelStar, "\nHotel Amenities: "+h.HotelAmenities, "\nHotel Price: ", h.Price, "\nCountry: "+h.Country)
		}

		defer db.Close()

	}
}

func getHotelBw300and400() {
	var h Hotel

	db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/ETIASSG2_db")
	if err != nil {
		panic(err.Error())

	}
	result1 := db.QueryRow("select * from Hotels where price > 300 and price < 400")
	err1 := result1.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
	if err1 == sql.ErrNoRows {
		fmt.Println("Sorry we do not have hotels from this price range")
		BookingMenu()
	} else {

		results, err := db.Query("select * from Hotels where price > 300 and price < 400")
		if err != nil {
			fmt.Println("Error")
			panic(err.Error())

		}
		for results.Next() {

			err = results.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("\nHotel Name: "+h.HotelName, "\nHotel Information: "+h.HotelInfo, "\nHotel Address: "+h.HotelAddr, "\nHotel Star: ", h.HotelStar, "\nHotel Amenities: "+h.HotelAmenities, "\nHotel Price: ", h.Price, "\nCountry: "+h.Country)
		}

		defer db.Close()

	}
}

func getallHotels() {
	var h Hotel

	db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/ETIASSG2_db")
	if err != nil {
		panic(err.Error())

	}
	result1 := db.QueryRow("select * from Hotels")
	err1 := result1.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
	if err1 == sql.ErrNoRows {
		fmt.Println("Error")
		BookingMenu()
	} else {

		results, err := db.Query("select * from Hotels")
		if err != nil {
			fmt.Println("Error")
			panic(err.Error())

		}
		for results.Next() {

			err = results.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("\nHotel Name: "+h.HotelName, "\nHotel Information: "+h.HotelInfo, "\nHotel Address: "+h.HotelAddr, "\nHotel Star: ", h.HotelStar, "\nHotel Amenities: "+h.HotelAmenities, "\nHotel Price: ", h.Price, "\nCountry: "+h.Country)
		}

		defer db.Close()

	}
}

func BookingMenu() {

	fmt.Println("Welcome, what is your name?")
	username = getInput()

	printTitle()

	for {
		printBookingMenu()
		text := getInput()

		var err error
		switch text {
		case "1":
			err = listAllPackages()
			printTitle()
		case "2":
			err = bookPackage()
			printTitle()
		case "3":
			err = changeBooking()
			printTitle()
		case "main":
			printTitle()
			continue
		case "quit", "exit":
			fmt.Println("Bye bye!")
			fmt.Println()
			main()
		default:
			printTitle()
			printStatusMessage(fmt.Sprintf("'%s' is not a correct option, try again", text))
			continue
		}

		if err != nil {
			printTitle()
			printErrStatus(err)
		}
	}

}

func listAllPackages() error {
	printTitle()
	printSubtitle("List of all hotels and attractions")
	_, err := getAndListPackages()
	if err != nil {
		return err
	}
	fmt.Println()
	enterToContinue()
	return nil
}

func bookPackage() error {
	printTitle()
	printSubtitle("Book a hotel or attraction")

	for {
		packages, err := getAndListPackages()
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Println(" Which package you want?")
		text := getInput()

		switch text {
		case "main":
			return nil
		case "quit", "exit":
			fmt.Println(" Bye bye!")
			fmt.Println()
			os.Exit(0)
		default:
			id, parseErr := strconv.Atoi(text)
			if parseErr != nil || id < 0 || id > len(packages) {
				printTitle()
				printSubtitle("Book a hotel or attraction")
				printStatusMessage(fmt.Sprintf("'%s' is not a correct option, try again", text))
				continue
			}

			index := id - 1
			err = performBooking(packages[index], username)
			if err != nil {
				printTitle()
				printSubtitle("Book a hotel or attraction")
				printErrStatus(err)
				continue
			}

			printTitle()
			printSubtitle("Book a hotel or attraction")
			printStatusMessage(fmt.Sprintf("Booking success! Booked package %s", packages[index].Name))
			enterToContinue()
			return nil
		}
	}
}

func getAndListPackages() ([]*PackageModel, error) {
	resp, err := http.Get("http://localhost:3000/packages")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var packages []*PackageModel

	err = json.Unmarshal(body, &packages)
	if err != nil {
		return nil, err
	}

	for idx, p := range packages {
		fmt.Printf(" %d. %s (type: %s) -- $%d\n", idx+1, p.Name, strings.ToUpper(p.PackageType), p.Price)
	}

	return packages, nil
}

func performBooking(packageModel *PackageModel, username string) error {
	url := fmt.Sprintf("http://localhost:3000/bookings?package_id=%d&username=%s", packageModel.ID, username)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	return nil
}

func changeBooking() error {
	printTitle()
	printSubtitle(fmt.Sprintf("List of bookings for '%s'", username))

	for {
		bookings, err := getAndListBookings(username)
		if err != nil {
			return err
		}

		if len(bookings) <= 0 {
			printStatusMessage("No bookings found")
			enterToContinue()
			return nil
		}

		fmt.Println()
		fmt.Println(" Which booking you want to change?")
		fmt.Println(" To cancel a booking, type: 'cancel <id>'")
		fmt.Println(" To update a booking, type: 'update <id>'")
		text := getInput()

		switch text {
		case "main":
			return nil
		case "quit", "exit":
			fmt.Println(" Bye bye!")
			fmt.Println()
			os.Exit(0)
		default:
			arr := strings.Split(text, " ")
			if len(arr) != 2 {
				printTitle()
				printSubtitle(fmt.Sprintf("List of bookings for '%s'", username))
				printStatusMessage(fmt.Sprintf("'%s' is not a correct option, try again", text))
				continue
			}

			command := arr[0]
			idString := arr[1]

			if command != "cancel" && command != "update" {
				printTitle()
				printSubtitle(fmt.Sprintf("List of bookings for '%s'", username))
				printStatusMessage(fmt.Sprintf("'%s' is not a correct option, try again", text))
				continue
			}

			id, parseErr := strconv.Atoi(idString)
			if parseErr != nil || id < 0 || id > len(bookings) {
				printTitle()
				printSubtitle(fmt.Sprintf("List of bookings for '%s'", username))
				printStatusMessage(fmt.Sprintf("'%s' is not a correct option, try again", text))
				continue
			}

			index := id - 1
			booking := bookings[index]
			if command == "cancel" {
				err = cancelBooking(booking)
				if err == nil {
					return nil
				}
			} else if command == "update" {
				err = updateBooking(booking)
				if err == nil {
					return nil
				}
			} else {
				panic("should not happen, command must be cancel or update!")
			}

			if err != nil {
				printTitle()
				printSubtitle(fmt.Sprintf("List of bookings for '%s'", username))
				printErrStatus(err)
				continue
			}
		}

		if err == nil {
			fmt.Println("Change of booking success!")
			fmt.Println("Press 'Enter' to return to main menu")
			_, _ = reader.ReadString('\n')
			return nil
		}
	}

}

func getAndListBookings(username string) ([]*BookedPackage, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/bookings?username=%s", username))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var bookedPackages []*BookedPackage

	err = json.Unmarshal(body, &bookedPackages)
	if err != nil {
		return nil, err
	}

	for idx, p := range bookedPackages {
		fmt.Printf(" %d. %s (type: %s) -- $%d\n", idx+1, p.PackageName, strings.ToUpper(p.PackageType), p.PackagePrice)
	}

	return bookedPackages, nil
}

func cancelBooking(booking *BookedPackage) error {
	url := fmt.Sprintf("http://localhost:3000/bookings?id=%d", booking.BookingId)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	printTitle()
	printSubtitle(fmt.Sprintf("List of bookings for '%s'", username))
	printStatusMessage("Successfully canceled booking!")
	enterToContinue()
	return nil
}

func updateBooking(booking *BookedPackage) error {
	printTitle()
	printSubtitle(fmt.Sprintf("Update booking '%s' for '%s'", booking.PackageName, username))

	for {
		fmt.Println(" List of packages:")
		packages, err := getAndListPackages()
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Println(" Which package you want to change to?")
		text := getInput()

		switch text {
		case "main":
			return nil
		case "quit", "exit":
			fmt.Println(" Bye bye!")
			fmt.Println()
			os.Exit(0)
		default:
			id, parseErr := strconv.Atoi(text)
			if parseErr != nil || id < 0 || id > len(packages) {
				printTitle()
				printSubtitle(fmt.Sprintf("Update booking '%s' for '%s'", booking.PackageName, username))
				printStatusMessage(fmt.Sprintf("'%s' is not a correct option, try again", text))
				continue
			}

			index := id - 1
			packageToChange := packages[index]
			url := fmt.Sprintf("http://localhost:3000/bookings?booking_id=%d&package_id=%d", booking.BookingId, packageToChange.ID)
			req, err := http.NewRequest("PATCH", url, nil)
			if err != nil {
				return err
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			if resp.StatusCode != 200 {
				return errors.New(resp.Status)
			}

			printTitle()
			printSubtitle(fmt.Sprintf("List of bookings for '%s'", username))
			printStatusMessage("Successfully updated booking!")
			enterToContinue()

			return nil
		}
	}
}
