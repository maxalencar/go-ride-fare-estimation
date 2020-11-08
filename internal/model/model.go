package model

import "time"

// Data represents the file data
type Data struct {
	RideID    int
	Latitude  float64
	Longitude float64
	Timestamp time.Time
}

// Ride represents the ride model containing the positions of a ride and the fare estimation
type Ride struct {
	ID           int
	Positions    []Position
	Segments     []Segment
	FareEstimate float64
}

// Position represents the position of the taxi
type Position struct {
	Coordinate Coordinate
	Timestamp  time.Time
}

// Coordinate represents a geographic coordinate.
type Coordinate struct {
	Latitude  float64
	Longitude float64
}

//Segment represents a segment between two positions.
type Segment struct {
	Position1 Position
	Position2 Position
	Distance  float64
	Duration  time.Duration
	Speed     float64
}
