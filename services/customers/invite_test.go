package customers

import (
	"encoding/json"
	"testing"

	"github.com/nbio/st"
)

//
// TestGetCustomersToInvite tests getting all customers within range, and sorted.
//
func TestGetCustomersToInvite(t *testing.T) {
	testFile := "../../data/customers.txt"

	customers, err := GetCustomersToInvite(testFile)
	st.Expect(t, err, nil)
	st.Expect(t, len(customers), 16)

	// We do not need to over test because of all our other testing. This is just a quick and dirty test
	json, _ := json.Marshal(customers)
	strTest := `[{"name":"Ian Kehoe","user_id":4,"latitude":"53.2451022","longitude":"-6.238335"},{"name":"Nora Dempsey","user_id":5,"latitude":"53.1302756","longitude":"-6.2397222"},{"name":"Theresa Enright","user_id":6,"latitude":"53.1229599","longitude":"-6.2705202"},{"name":"Eoin Ahearn","user_id":8,"latitude":"54.0894797","longitude":"-6.18671"},{"name":"Richard Finnegan","user_id":11,"latitude":"53.008769","longitude":"-6.1056711"},{"name":"Christina McArdle","user_id":12,"latitude":"52.986375","longitude":"-6.043701"},{"name":"Olive Ahearn","user_id":13,"latitude":"53","longitude":"-7"},{"name":"Michael Ahearn","user_id":15,"latitude":"52.966","longitude":"-6.463"},{"name":"Patricia Cahill","user_id":17,"latitude":"54.180238","longitude":"-5.920898"},{"name":"Eoin Gallagher","user_id":23,"latitude":"54.080556","longitude":"-6.361944"},{"name":"Rose Enright","user_id":24,"latitude":"54.133333","longitude":"-6.433333"},{"name":"Stephen McArdle","user_id":26,"latitude":"53.038056","longitude":"-7.653889"},{"name":"Oliver Ahearn","user_id":29,"latitude":"53.74452","longitude":"-7.11167"},{"name":"Nick Enright","user_id":30,"latitude":"53.761389","longitude":"-7.2875"},{"name":"Alan Behan","user_id":31,"latitude":"53.1489345","longitude":"-6.8422408"},{"name":"Lisa Ahearn","user_id":39,"latitude":"53.0033946","longitude":"-6.3877505"}]`
	st.Expect(t, string(json), strTest)
}

//
// TestSortCustomersById will test sorting customers by ID
//
func TestSortCustomersById(t *testing.T) {
	testFile := "../../data/customers.txt"

	customers, err := readCustomerList(testFile)
	st.Expect(t, err, nil)

	sortCustomersById(customers)

	var lastID int64 = 0

	for _, row := range customers {
		if lastID == 0 {
			lastID = row.UserID
			continue
		}

		st.Expect(t, row.UserID > lastID, true)

		lastID = row.UserID
	}
}

//
// TestGetCustomersWithinRange will test returning customers with in a certain range of the office
//
func TestGetCustomersWithinRange(t *testing.T) {
	testFile := "../../data/customers.txt"

	customers, err := readCustomerList(testFile)
	st.Expect(t, err, nil)
	st.Expect(t, len(customers), 32)

	filteredCustomers, err := getCustomersWithinRange(customers)
	st.Expect(t, err, nil)
	st.Expect(t, len(filteredCustomers), 16)

	// Spot check the different rows.
	st.Expect(t, filteredCustomers[0].Name, "Christina McArdle")
	st.Expect(t, filteredCustomers[0].UserID, int64(12))
	st.Expect(t, filteredCustomers[0].Latitude, 52.986375)
	st.Expect(t, filteredCustomers[0].Longitude, -6.043701)
	st.Expect(t, filteredCustomers[5].Name, "Nora Dempsey")
	st.Expect(t, filteredCustomers[5].UserID, int64(5))
	st.Expect(t, filteredCustomers[5].Latitude, 53.1302756)
	st.Expect(t, filteredCustomers[5].Longitude, -6.2397222)
	st.Expect(t, filteredCustomers[15].Name, "Eoin Gallagher")
	st.Expect(t, filteredCustomers[15].UserID, int64(23))
	st.Expect(t, filteredCustomers[15].Latitude, 54.080556)
	st.Expect(t, filteredCustomers[15].Longitude, -6.361944)

	// One last check to make sure we did not miss anything in our spot checking above.
	json, _ := json.Marshal(filteredCustomers)
	strTest := `[{"name":"Christina McArdle","user_id":12,"latitude":"52.986375","longitude":"-6.043701"},{"name":"Eoin Ahearn","user_id":8,"latitude":"54.0894797","longitude":"-6.18671"},{"name":"Stephen McArdle","user_id":26,"latitude":"53.038056","longitude":"-7.653889"},{"name":"Theresa Enright","user_id":6,"latitude":"53.1229599","longitude":"-6.2705202"},{"name":"Ian Kehoe","user_id":4,"latitude":"53.2451022","longitude":"-6.238335"},{"name":"Nora Dempsey","user_id":5,"latitude":"53.1302756","longitude":"-6.2397222"},{"name":"Richard Finnegan","user_id":11,"latitude":"53.008769","longitude":"-6.1056711"},{"name":"Alan Behan","user_id":31,"latitude":"53.1489345","longitude":"-6.8422408"},{"name":"Olive Ahearn","user_id":13,"latitude":"53","longitude":"-7"},{"name":"Michael Ahearn","user_id":15,"latitude":"52.966","longitude":"-6.463"},{"name":"Patricia Cahill","user_id":17,"latitude":"54.180238","longitude":"-5.920898"},{"name":"Lisa Ahearn","user_id":39,"latitude":"53.0033946","longitude":"-6.3877505"},{"name":"Rose Enright","user_id":24,"latitude":"54.133333","longitude":"-6.433333"},{"name":"Oliver Ahearn","user_id":29,"latitude":"53.74452","longitude":"-7.11167"},{"name":"Nick Enright","user_id":30,"latitude":"53.761389","longitude":"-7.2875"},{"name":"Eoin Gallagher","user_id":23,"latitude":"54.080556","longitude":"-6.361944"}]`
	st.Expect(t, string(json), strTest)
}

//
// TestReadCustomerList will test reading in customers and converting them to an array of structs
//
func TestReadCustomerList(t *testing.T) {
	testFile := "../../data/customers.txt"

	rows, err := readCustomerList(testFile)
	st.Expect(t, err, nil)

	// Spot check the different rows.
	st.Expect(t, rows[0].Name, "Christina McArdle")
	st.Expect(t, rows[0].UserID, int64(12))
	st.Expect(t, rows[0].Latitude, 52.986375)
	st.Expect(t, rows[0].Longitude, -6.043701)
	st.Expect(t, rows[15].Name, "Alan Behan")
	st.Expect(t, rows[15].UserID, int64(31))
	st.Expect(t, rows[15].Latitude, 53.1489345)
	st.Expect(t, rows[15].Longitude, -6.8422408)
	st.Expect(t, rows[31].Name, "David Behan")
	st.Expect(t, rows[31].UserID, int64(25))
	st.Expect(t, rows[31].Latitude, 52.833502)
	st.Expect(t, rows[31].Longitude, -8.522366)

	// One last check to make sure we did not miss anything in our spot checking above.
	json, _ := json.Marshal(rows)
	strTest := `[{"name":"Christina McArdle","user_id":12,"latitude":"52.986375","longitude":"-6.043701"},{"name":"Alice Cahill","user_id":1,"latitude":"51.92893","longitude":"-10.27699"},{"name":"Ian McArdle","user_id":2,"latitude":"51.8856167","longitude":"-10.4240951"},{"name":"Jack Enright","user_id":3,"latitude":"52.3191841","longitude":"-8.5072391"},{"name":"Charlie Halligan","user_id":28,"latitude":"53.807778","longitude":"-7.714444"},{"name":"Frank Kehoe","user_id":7,"latitude":"53.4692815","longitude":"-9.436036"},{"name":"Eoin Ahearn","user_id":8,"latitude":"54.0894797","longitude":"-6.18671"},{"name":"Stephen McArdle","user_id":26,"latitude":"53.038056","longitude":"-7.653889"},{"name":"Enid Gallagher","user_id":27,"latitude":"54.1225","longitude":"-8.143333"},{"name":"Theresa Enright","user_id":6,"latitude":"53.1229599","longitude":"-6.2705202"},{"name":"Jack Dempsey","user_id":9,"latitude":"52.2559432","longitude":"-7.1048927"},{"name":"Georgina Gallagher","user_id":10,"latitude":"52.240382","longitude":"-6.972413"},{"name":"Ian Kehoe","user_id":4,"latitude":"53.2451022","longitude":"-6.238335"},{"name":"Nora Dempsey","user_id":5,"latitude":"53.1302756","longitude":"-6.2397222"},{"name":"Richard Finnegan","user_id":11,"latitude":"53.008769","longitude":"-6.1056711"},{"name":"Alan Behan","user_id":31,"latitude":"53.1489345","longitude":"-6.8422408"},{"name":"Olive Ahearn","user_id":13,"latitude":"53","longitude":"-7"},{"name":"Helen Cahill","user_id":14,"latitude":"51.999447","longitude":"-9.742744"},{"name":"Michael Ahearn","user_id":15,"latitude":"52.966","longitude":"-6.463"},{"name":"Ian Larkin","user_id":16,"latitude":"52.366037","longitude":"-8.179118"},{"name":"Patricia Cahill","user_id":17,"latitude":"54.180238","longitude":"-5.920898"},{"name":"Lisa Ahearn","user_id":39,"latitude":"53.0033946","longitude":"-6.3877505"},{"name":"Bob Larkin","user_id":18,"latitude":"52.228056","longitude":"-7.915833"},{"name":"Rose Enright","user_id":24,"latitude":"54.133333","longitude":"-6.433333"},{"name":"Enid Cahill","user_id":19,"latitude":"55.033","longitude":"-8.112"},{"name":"Enid Enright","user_id":20,"latitude":"53.521111","longitude":"-9.831111"},{"name":"David Ahearn","user_id":21,"latitude":"51.802","longitude":"-9.442"},{"name":"Charlie McArdle","user_id":22,"latitude":"54.374208","longitude":"-8.371639"},{"name":"Oliver Ahearn","user_id":29,"latitude":"53.74452","longitude":"-7.11167"},{"name":"Nick Enright","user_id":30,"latitude":"53.761389","longitude":"-7.2875"},{"name":"Eoin Gallagher","user_id":23,"latitude":"54.080556","longitude":"-6.361944"},{"name":"David Behan","user_id":25,"latitude":"52.833502","longitude":"-8.522366"}]`
	st.Expect(t, string(json), strTest)
}

//
// TestValidateCustomerList - test uploading a customer list
//
func TestValidateCustomerList(t *testing.T) {
	testFile1 := "../../data/customers.txt"
	testFile2 := "../../data/customers.json"

	err := ValidateCustomerList(testFile1)
	st.Expect(t, err, nil)

	err = ValidateCustomerList(testFile2)
	st.Expect(t, err.Error(), "uploaded: invalid file type")

}
