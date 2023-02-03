package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Itinerarie struct {
	Location    string `json:"Location"`
	DurOfTravel int    `json:"Duration"`
	StartDate   string `json:"Start Date"`
	EndDate     string `json:"End Date"`
}

type AllItineraries struct {
	Itineraries map[string]Itinerarie `json:"Itineraries"`
}

var itineraries map[string]Itinerarie = map[string]Itinerarie{
	" ": Itinerarie{"MBS", 8, "03/02/2023", "05/02/2023"},
}

type Test struct {
	Location string `json:"Location"`
	Value    int    `json:"Value"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/itineraries", itinerariesFilter)
	router.HandleFunc("/api/v1/itineraries/{itineraries_id}", allitineraries).Methods("GET", "POST", "PUT")

	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func itinerariesFilter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	results := map[string]Itinerarie{}
	if value := query.Get("q"); len(value) > 0 {
		for k, v := range itineraries {
			if strings.Contains(strings.ToLower(v.Location), strings.ToLower(value)) {
				results[k] = v
			}
		}

		if len(results) == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No itinerarie found")
		} else {
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Itinerarie `json:"Search Results"`
			}{results})
		}
	} else if value := query.Get("value"); len(value) > 0 {
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body), "\n", err)

		for k, v := range itineraries {
			value, _ := strconv.Atoi(value)
			if v.DurOfTravel >= value {
				results[k] = v
			}
		}

		if len(results) == 0 {
			fmt.Fprintf(w, "No itinerarie eligible")
		} else {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(struct {
				SearchResults map[string]Itinerarie `json:"Eligible Itinerarie(s)"`
			}{results})
		}
	} else {
		allitineraries := AllItineraries{itineraries}

		json.NewEncoder(w).Encode(allitineraries)
	}
}

func allitineraries(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Println(params["itineraries_id"], r.Method)

	if v, ok := itineraries[params["itineraries_id"]]; ok {
		if r.Method == "GET" {
			json.NewEncoder(w).Encode(v)
		} else if r.Method == "POST" {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "Itinerarie ID exists")
		} else if r.Method == "PUT" {
			if body, err := ioutil.ReadAll(r.Body); err == nil {
				var data Itinerarie

				if err := json.Unmarshal(body, &data); err == nil {
					fmt.Printf("PUT ### %v", data)
					w.WriteHeader(http.StatusAccepted)
					itineraries[params["itineraries_id"]] = data
				}
			}
		} else {
			delete(itineraries, params["itineraries_id"])
			fmt.Fprintf(w, params["itineraries_id"]+" Deleted")
		}
	} else if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Itinerarie

			if err := json.Unmarshal(body, &data); err == nil {
				w.WriteHeader(http.StatusAccepted)
				itineraries[params["itineraries_id"]] = data
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		if r.Method == "PUT" {
			fmt.Fprintf(w, "Itinerarie ID does not exist")
		} else {
			fmt.Fprintf(w, "Invalid Itinerarie ID")
		}
	}
}
