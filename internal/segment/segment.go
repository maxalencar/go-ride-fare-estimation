package segment

import (
	"go-ride-fare-estimation/internal/harvesine"
	"go-ride-fare-estimation/internal/model"
)

const (
	speedLimit = 100 // speed limit to be considered a valid entry.
)

// Create a list of valid segments of a ride based on the positions
// 1. it considers a segment as invalid if the speed is more than 100 km/h
func Create(positions []model.Position) []model.Segment {
	var segments = make([]model.Segment, 0)

	for i := 0; i < len(positions)-1; {
		nextPos := 1

		var s model.Segment

		for i+nextPos < len(positions) {
			s = calculate(positions[i], positions[i+nextPos])

			// if the segment speed is higher than the speed limit, it is an invalid entry
			if s.Speed > speedLimit {
				nextPos++
				continue
			}

			break
		}

		// it skips invalid entries
		i += nextPos

		// in case the last segment is invalid, we don't add it to the list
		if s.Speed <= speedLimit {
			segments = append(segments, s)
		}
	}

	return segments
}

// calculate a segment based on two positions
// and return the distance in km,
// duration and the speed in km/h
func calculate(p1, p2 model.Position) model.Segment {
	km, _ := harvesine.Distance(p1.Coordinate, p2.Coordinate)
	duration := p2.Timestamp.Sub(p1.Timestamp)
	speed := km / duration.Hours()

	return model.Segment{
		Position1: p1,
		Position2: p2,
		Distance:  km,
		Duration:  duration,
		Speed:     speed,
	}
}
