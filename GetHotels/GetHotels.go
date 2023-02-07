package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

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

type AllHotels struct {
	Hotels map[string]Hotel `json:"Hotels"`
}

var hotels map[string]Hotel = map[string]Hotel{
	" ": Hotel{1, "PARKROYAL", "Polished rooms in a chic hotel featuring an outdoor pool 2 restaurants & a gym", "181 Kitchener Rd, Singapore 208533", 4, "Pool,Parking,WIFI,Air Conditioning", 244, "Singapore"},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/hotels/Country", CountryFilter)
	router.HandleFunc("/api/v1/hotels/HotelStar", HotelStarFilter)
	router.HandleFunc("/api/v1/hotels/Amenities", AmenitiesFilter)
	router.HandleFunc("/api/v1/hotels/Price", PriceFilter)
	router.HandleFunc("/api/v1/hotels/{hotel_id}", allhotels).Methods("GET")

	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func PriceFilter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	results := map[string]Hotel{}
	if value := query.Get("q"); len(value) > 0 {
		for k, v := range hotels {
			results[k] = v
		}

		if len(results) == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No Hotel found")
		} else {
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Hotel `json:"Search Results"`
			}{results})
		}
	} else if value := query.Get("value"); len(value) > 0 {
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body), "\n", err)

		if len(results) == 0 {
			fmt.Fprintf(w, "No Hotel Found")
		} else {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Hotel `json:"Hotels Found"`
			}{results})
		}
	} else {
		allhotels := AllHotels{hotels}

		json.NewEncoder(w).Encode(allhotels)
	}
}

func HotelStarFilter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	results := map[string]Hotel{}
	if value := query.Get("q"); len(value) > 0 {
		for k, v := range hotels {
			results[k] = v
		}

		if len(results) == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No Hotel found")
		} else {
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Hotel `json:"Search Results"`
			}{results})
		}
	} else if value := query.Get("value"); len(value) > 0 {
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body), "\n", err)

		if len(results) == 0 {
			fmt.Fprintf(w, "No Hotel Found")
		} else {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Hotel `json:"Hotels Found"`
			}{results})
		}
	} else {
		allhotels := AllHotels{hotels}

		json.NewEncoder(w).Encode(allhotels)
	}
}

func AmenitiesFilter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	results := map[string]Hotel{}
	if value := query.Get("q"); len(value) > 0 {
		for k, v := range hotels {
			if strings.Contains(strings.ToLower(v.HotelAmenities), strings.ToLower(value)) {
				results[k] = v
			}
		}

		if len(results) == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No Hotel found")
		} else {
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Hotel `json:"Search Results"`
			}{results})
		}
	} else if value := query.Get("value"); len(value) > 0 {
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body), "\n", err)

		if len(results) == 0 {
			fmt.Fprintf(w, "No Hotel Found")
		} else {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Hotel `json:"Hotels Found"`
			}{results})
		}
	} else {
		allhotels := AllHotels{hotels}

		json.NewEncoder(w).Encode(allhotels)
	}
}

func CountryFilter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	results := map[string]Hotel{}
	if value := query.Get("q"); len(value) > 0 {
		for k, v := range hotels {
			if strings.Contains(strings.ToLower(v.Country), strings.ToLower(value)) {
				results[k] = v
			}
		}

		if len(results) == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No Hotel found")
		} else {
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Hotel `json:"Search Results"`
			}{results})
		}
	} else if value := query.Get("value"); len(value) > 0 {
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body), "\n", err)

		if len(results) == 0 {
			fmt.Fprintf(w, "No Hotel Found")
		} else {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Hotel `json:"Hotels Found"`
			}{results})
		}
	} else {
		allhotels := AllHotels{hotels}

		json.NewEncoder(w).Encode(allhotels)
	}
}

func allhotels(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Println(params["hotel_id"], r.Method)

	if v, ok := hotels[params["hotel_id"]]; ok {
		if r.Method == "GET" {
			json.NewEncoder(w).Encode(v)
		}
	}
}
