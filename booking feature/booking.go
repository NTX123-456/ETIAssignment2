package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

/** Start: Models **/

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

func printMenu() {
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

func main() {

	fmt.Println("Welcome, what is your name?")
	username = getInput()

	printTitle()

	for {
		printMenu()
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
			fmt.Println(" Bye bye!")
			fmt.Println()
			os.Exit(0)
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
