package geo

import "math"

//
// CalDistance will take a starting lat / long, and a ending lat / long and return the distance in KM.
// Used the math highlighted here https://en.wikipedia.org/wiki/Great-circle_distance
// Modeled code after https://www.geodatasource.com/developers/go
//
func CalDistance(latStart float64, longStart float64, latEnd float64, longEnd float64) float64 {
	const PI float64 = 3.141592653589793

	radLatStart := float64(PI * latStart / 180)
	radLatLong := float64(PI * latEnd / 180)
	theta := float64(longStart - longEnd)
	radTheta := float64(PI * theta / 180)

	distance := math.Sin(radLatStart)*math.Sin(radLatLong) + math.Cos(radLatStart)*math.Cos(radLatLong)*math.Cos(radTheta)

	if distance > 1 {
		distance = 1
	}

	distance = math.Acos(distance)
	distance = distance * 180 / PI
	distance = distance * 60 * 1.1515

	// Convert to KM
	distance = distance * 1.609344

	// Return a rounded result
	return math.Round(distance*100) / 100
}
