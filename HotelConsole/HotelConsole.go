package main

import (
	//"bufio"
	//"bytes"
	"database/sql"
	"fmt"

	"log"
	"net/http"
	"os"
	"strings"

	//"io/ioutil"
	//"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

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

func Hotelmain() {
outer:
	for {
		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("You are Now In the Hotel Console")
		fmt.Println("Get Hotels\n",
			"1. List all Hotels\n",
			"2. List Hotels by Country\n",
			"3. List Hotels by Hotel Star\n",
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

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Welcome to the Hotel Console\n</h1>")

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
		main()
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

	fmt.Println("Enter the amount of stars of the hotel: ")
	fmt.Scanf("%d\n", &h.HotelStar)

	db, err := sql.Open("mysql", "root:Lolman@4567@tcp(127.0.0.1:3306)/ETIASSG2_db")
	if err != nil {
		panic(err.Error())

	}
	result1 := db.QueryRow("select * from Hotels where HotelStar = ?", &h.HotelStar)
	err1 := result1.Scan(&h.ID, &h.HotelName, &h.HotelInfo, &h.HotelAddr, &h.HotelStar, &h.HotelAmenities, &h.Price, &h.Country)
	if err1 == sql.ErrNoRows {
		fmt.Println("Sorry we do not have hotels from this star")
		main()
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
		main()
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
		main()
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
		main()
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
		main()
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
		main()
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
		main()
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
		main()
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
