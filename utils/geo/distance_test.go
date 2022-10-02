package geo

import (
	"testing"

	"github.com/nbio/st"
)

//
// TestReadCustomerList will test reading in customers and converting them to an array of structs
// Verified results using this: https://boulter.com/gps/distance/?from=45.3064755%2C+-122.9751364&to=45.3058547%2C+-122.9456318&units=k
//
func TestCalDistance(t *testing.T) {
	dist := CalDistance(32.9697, -96.80322, 29.46786, -98.53506)
	st.Expect(t, dist, 422.74)

	dist = CalDistance(43.9597048, -75.9040299, 45.3064755, -122.9751364)
	st.Expect(t, dist, 3673.93)

	dist = CalDistance(45.3064755, -122.9751364, 45.3058547, -122.9456318)
	st.Expect(t, dist, 2.31)
}
