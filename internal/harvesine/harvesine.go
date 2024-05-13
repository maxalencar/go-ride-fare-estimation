package harvesine

import (
	"math"

	"go-ride-fare-estimation/internal/model"
)

const (
	earthRaidusKm = 6371 // radius of the earth in kilometers.
	earthRadiusMi = 3958 // radius of the earth in miles.
)

// Distance calculates the shortest path between two coordinates on the surface
// of the Earth. This function returns two units of measure, the first is the
// distance in kilometers and the second is the distance in miles.
func Distance(c1, c2 model.Coordinate) (km, mi float64) {
	lat1 := degreesToRadians(c1.Latitude)
	long1 := degreesToRadians(c1.Longitude)
	lat2 := degreesToRadians(c2.Latitude)
	long2 := degreesToRadians(c2.Longitude)

	diffLat := lat2 - lat1
	diffLong := long2 - long1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLong/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	km = c * earthRaidusKm
	mi = c * earthRadiusMi

	return km, mi
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}
