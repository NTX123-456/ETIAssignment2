package main

import (
	//"bufio"
	//"bytes"
	//"encoding/json"
	"database/sql"
	"fmt"

	//"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
outer:
	for {
		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("Get Hotels\n",
			"1. List all Hotels\n",
			"2. List all Hotels\n",
			"3. Quit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			getallHotels()
		case 2:
			getHotelsbyCountry()
		case 3:
			break outer
		}
	}
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
