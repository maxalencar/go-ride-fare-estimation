package fare

import (
	"time"

	"go-ride-fare-estimation/internal/model"
)

const (
	speedMoving = 10 // speed that is considered for the taxi being in moving state
)

const (
	priceNormalFare  = 0.74  // price per km
	priceExtraFare   = 1.30  // price per km
	priceInitialFare = 1.30  // initial fare price
	priceMinimumFare = 3.47  // minimum fare price
	priceIdleFare    = 11.90 // price per hour of idle time
)

// Calculate calculates the fare of a ride
func Calculate(segments []model.Segment) float64 {
	return calculateFareEstimate(segments)
}

// calculateFareEstimate calculates the fare estimate based on the segments
// 1. it starts the fare estimate with the initial fare price of 1.30
// 2. it applies the price based on the time/state of the segment
// 3. it applies the minimum fare if the total is less than 3.47
func calculateFareEstimate(segments []model.Segment) float64 {
	fareEstimate := priceInitialFare

	for _, v := range segments {
		if v.Speed > speedMoving {
			if isNormalFare(v.Position1.Timestamp) {
				fareEstimate += v.Distance * priceNormalFare
			} else {
				fareEstimate += v.Distance * priceExtraFare
			}
		} else {
			fareEstimate += v.Duration.Hours() * priceIdleFare
		}
	}

	if fareEstimate < priceMinimumFare {
		fareEstimate = priceMinimumFare
	}

	return fareEstimate
}

// isNormalFare checks if the segment is between the time considered a normal fare.
// Time of day [05:00, 00:00]
// Assuming that 00:00 will be 23:59 as 00:00 is already the next day
// so we consider from 00:00 as extra
func isNormalFare(timestamp time.Time) bool {
	return timestamp.Hour() >= 5 && timestamp.Hour() <= 23
}
