package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/driver/mysql"
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

/** Start: DB code **/

var db *gorm.DB

const (
	dbUsername = "user"
	dbPassword = "password"
	dbAddress  = "127.0.0.1"
	dbPort     = "3306"
	dbName     = "booking_db"
)

func initDatabase() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbAddress, dbPort, dbName)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to init db with err: %v", err))
	}

	err = db.AutoMigrate(&PackageModel{}, &BookingModel{})
	if err != nil {
		panic(fmt.Sprintf("failed to auto migrate db with err: %v", err))
	}

	fmt.Println("successfully init db")
}

func getPackages() ([]*PackageModel, error) {
	var packages []*PackageModel
	err := db.Debug().Find(&packages).Error
	if err != nil {
		fmt.Printf("error while finding packages, err: %v\n", err)
		return nil, err
	}
	return packages, nil
}

func insertPackage(p *PackageModel) error {
	err := db.Debug().Create(p).Error
	if err != nil {
		fmt.Printf("error while inserting package, err: %v\n", err)
		return err
	}
	return nil
}

func deletePackageById(id int) error {
	err := db.Debug().Where("id = ?", id).Delete(&PackageModel{}).Error
	if err != nil {
		fmt.Printf("error while deleting package, err: %v\n", err)
		return err
	}
	return nil
}

func getBookings(username string) ([]*BookingModel, error) {
	var bookings []*BookingModel
	err := db.Debug().Where("username = ?", username).Find(&bookings).Error
	if err != nil {
		fmt.Printf("error while finding bookings, err: %v\n", err)
		return nil, err
	}
	return bookings, nil
}

func insertBooking(b *BookingModel) error {
	err := db.Debug().Create(b).Error
	if err != nil {
		fmt.Printf("error while inserting booking, err: %v\n", err)
		return err
	}
	return nil
}

func deleteBookingById(id int) error {
	err := db.Debug().Where("id = ?", id).Delete(&BookingModel{}).Error
	if err != nil {
		fmt.Printf("error while deleting booking, err: %v\n", err)
		return err
	}
	return nil
}

func updateBookingPackage(bookingId, packageId int) error {
	err := db.Debug().Where("id = ?", bookingId).Model(&BookingModel{}).UpdateColumn("package_id", packageId).Error
	if err != nil {
		fmt.Printf("error while deleting booking, err: %v\n", err)
		return err
	}
	return nil
}

/** End: DB code **/

/** Start: HTTP server code **/

func packagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		packages, err := getPackages()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		j, _ := json.Marshal(packages)
		_, err = w.Write(j)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "POST":
		d := json.NewDecoder(r.Body)
		p := &PackageModel{}
		err := d.Decode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = insertPackage(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "DELETE":
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = deletePackageById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "cannot", http.StatusMethodNotAllowed)
	}
}

func bookingsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		username := r.URL.Query().Get("username")
		bookings, err := getBookings(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packages, err := getPackages()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var bookedPackages []*BookedPackage
		for _, b := range bookings {
			for _, p := range packages {
				if b.PackageId == int(p.ID) {
					bookedPackages = append(bookedPackages, &BookedPackage{
						BookingId:       int(b.ID),
						BookingUsername: b.Username,
						PackageId:       b.PackageId,
						PackageName:     p.Name,
						PackageType:     p.PackageType,
						PackagePrice:    p.Price,
					})
				}
			}
		}

		j, _ := json.Marshal(bookedPackages)
		_, err = w.Write(j)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "POST":
		packageIdString := r.URL.Query().Get("package_id")
		packageId, err := strconv.Atoi(packageIdString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		username := r.URL.Query().Get("username")
		b := &BookingModel{
			PackageId: packageId,
			Username:  username,
		}
		err = insertBooking(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "PATCH":
		bookingIdString := r.URL.Query().Get("booking_id")
		bookingId, err := strconv.Atoi(bookingIdString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packageIdString := r.URL.Query().Get("package_id")
		packageId, err := strconv.Atoi(packageIdString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = updateBookingPackage(bookingId, packageId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case "DELETE":
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = deleteBookingById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "cannot", http.StatusMethodNotAllowed)
	}
}

func runServer() {
	http.HandleFunc("/packages", packagesHandler)
	http.HandleFunc("/bookings", bookingsHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(fmt.Sprintf("failed to run server with err: %v", err))
	}
}

/** End: HTTP server code **/

func main() {
	initDatabase()

	runServer()
}
