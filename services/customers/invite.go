package customers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"party_invite/utils/geo"
	"party_invite/utils/validation"
)

// TODO(spicer): Maybe these are better in a .env file or passed in via the web service or something.
const DistanceFromOffice float64 = 100 // KM
const OfficeLatitude float64 = 53.339428
const OfficeLongitude float64 = -6.257664

// Customer struct is an object to hold the information of one customer
type Customer struct {
	Name      string  `json:"name"`
	UserID    int64   `json:"user_id"`
	Latitude  float64 `json:"latitude,string"`
	Longitude float64 `json:"longitude,string"`
}

//
// GetCustomersToInvite will take a file path of a list of customers.
// We take that list and find the customer we want to invite based on
// how far away from our office they are.
//
func GetCustomersToInvite(uploadedFile string) ([]Customer, error) {
	// Validate if this is a valid file.
	err := ValidateCustomerList(uploadedFile)

	if err != nil {
		return []Customer{}, err
	}

	// Read the file into memory
	// TODO(spicer): we are assuming the file is big enough to fit into memory. Is this a bad assumption?
	customers, err := readCustomerList(uploadedFile)

	if err != nil {
		return []Customer{}, err
	}

	// Get customers within range
	nearCustomers, err := getCustomersWithinRange(customers)

	if err != nil {
		return []Customer{}, err
	}

	// Sort the near customers by UserID
	sortCustomersById(nearCustomers)

	return nearCustomers, nil
}

//
// ValidateCustomerList will make sure the file we uploaded is valid.
//
func ValidateCustomerList(uploadedFile string) error {
	// Validate the file.
	test, err := validation.FileIsTextFile(uploadedFile)

	if err != nil {
		return err
	}

	if !test {
		return fmt.Errorf("uploaded: invalid file type")
	}

	return nil
}

//
// sortCustomersById will take an array of customers and sort them in ascending order by user id.
// This function required Go 1.18+
//
func sortCustomersById(customers []Customer) {
	sort.Slice(customers, func(i, j int) bool {
		return customers[i].UserID < customers[j].UserID
	})
}

//
// getCustomersWithinRange will search for customers within a particular range
//
func getCustomersWithinRange(customers []Customer) ([]Customer, error) {
	rt := []Customer{}

	// Loop through the customers and capture the ones within DistanceFromOffice
	for _, row := range customers {
		dist := geo.CalDistance(OfficeLatitude, OfficeLongitude, row.Latitude, row.Longitude)

		if dist <= DistanceFromOffice {
			rt = append(rt, row)
		}
	}

	return rt, nil
}

//
// readCustomerList will read the uploaded file into memory and assign it to an array of customer stucts
//
func readCustomerList(uploadedFile string) ([]Customer, error) {
	rt := []Customer{}

	readFile, err := os.Open(uploadedFile)

	if err != nil {
		return rt, err
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	// Loop through each line and Unmarshal the json
	// TODO(spicer): Maybe consider modifying the text file to be proper JSON and then Unmarshal
	// the entire file instead of row by row. Likely not faster but we could test.
	for fileScanner.Scan() {
		row := Customer{}
		str := fileScanner.Text()

		if err := json.Unmarshal([]byte(str), &row); err != nil {
			return rt, err
		}

		rt = append(rt, row)
	}

	readFile.Close()

	return rt, nil
}
