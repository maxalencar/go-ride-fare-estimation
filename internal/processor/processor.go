package processor

import (
	"encoding/csv"
	"fmt"
	"go-ride-fare-estimation/internal/fare"
	"go-ride-fare-estimation/internal/model"
	"go-ride-fare-estimation/internal/segment"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

// Processor defines the minimum contract our processor must satisfy.
type Processor interface {
	CalculateFare(in <-chan model.Ride, wg *sync.WaitGroup) <-chan model.Ride
	CreateSegments(in <-chan model.Ride, wg *sync.WaitGroup) <-chan model.Ride
	Process(in <-chan model.Data, wg *sync.WaitGroup) <-chan model.Ride
	Read(filePath string, wg *sync.WaitGroup) <-chan model.Data
	WriteResult(in <-chan model.Ride, filePath string, wg *sync.WaitGroup)
}

// processor holds the structure of our processor implementation.
type processor struct {
}

// NewProcessor creates a new Processor.
func NewProcessor() Processor {
	return &processor{}
}

// CalculateFare calculates a fare based on the ride segments.
func (p processor) CalculateFare(in <-chan model.Ride, wg *sync.WaitGroup) <-chan model.Ride {
	wg.Add(1)
	out := make(chan model.Ride)

	go func() {
		for ride := range in {
			ride.FareEstimate = fare.Calculate(ride.Segments)
			out <- ride
		}

		close(out)
		wg.Done()
	}()

	return out
}

// CreateSegments creates a list of segments based on the ride positions
// ignoring invalid segments.
func (p processor) CreateSegments(in <-chan model.Ride, wg *sync.WaitGroup) <-chan model.Ride {
	wg.Add(1)
	out := make(chan model.Ride)

	go func() {
		for ride := range in {
			ride.Segments = segment.Create(ride.Positions)
			out <- ride
		}

		close(out)
		wg.Done()
	}()

	return out
}

// Process the data of a ride and send to a channel when all data of a ride is collected.
func (p processor) Process(in <-chan model.Data, wg *sync.WaitGroup) <-chan model.Ride {
	wg.Add(1)
	out := make(chan model.Ride)
	var r model.Ride

	go func() {
		for data := range in {
			if r.ID == 0 {
				r = model.Ride{ID: data.RideID}
			} else if r.ID != data.RideID {
				out <- r

				r = model.Ride{ID: data.RideID}
			}

			r.Positions = append(r.Positions, parseDataToPosition(data))
		}

		out <- r

		close(out)
		wg.Done()
	}()

	return out
}

// Read a file and convert the records into data struct.
func (p processor) Read(filePath string, wg *sync.WaitGroup) <-chan model.Data {
	wg.Add(1)
	out := make(chan model.Data)

	go func() {
		f, _ := os.Open(filePath)
		r := csv.NewReader(f)
		r.FieldsPerRecord = 4

		line := 0
		for {
			line++

			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			data, err := transform(record, line)
			if err != nil {
				log.Fatal(err)
			}

			out <- data
		}

		close(out)
		wg.Done()
	}()

	return out
}

// WriteResult creates a file with the result of the processing
// containing a list of ride_id and fare_estimate.
func (p processor) WriteResult(in <-chan model.Ride, filePath string, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("cannot create file; err: %v", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		for ride := range in {
			err := writer.Write([]string{strconv.Itoa(ride.ID), strconv.FormatFloat(ride.FareEstimate, 'f', 10, 64)})
			if err != nil {
				log.Fatalf("cannot write to file; err: %v", err)
			}
		}
	}()
}

// transform a record from the csv file into data struct.
func transform(record []string, line int) (model.Data, error) {
	var data model.Data

	// if the record contains less than four elements, it means a broken file so it returns an error
	if len(record) < 4 {
		return data, fmt.Errorf("invalid record. it should contains 4 elements, but it contains %d; line %d", len(record), line)
	}

	rideID, err := strconv.Atoi(record[0])
	if err != nil {
		return data, fmt.Errorf("could not parse timestamp %s; line %d; err: %v", record[3], line, err)
	}

	lat, err := strconv.ParseFloat(record[1], 64)
	if err != nil {
		return data, fmt.Errorf("could not parse latitute %s; line %d; err: %v", record[1], line, err)
	}

	long, err := strconv.ParseFloat(record[2], 64)
	if err != nil {
		return data, fmt.Errorf("could not parse longitude %s; line %d; err: %v", record[2], line, err)
	}

	timestamp, err := strconv.ParseInt(record[3], 10, 64)
	if err != nil {
		return data, fmt.Errorf("could not parse timestamp %s; line %d; err: %v", record[3], line, err)
	}

	data.RideID = rideID
	data.Latitude = lat
	data.Longitude = long
	data.Timestamp = time.Unix(timestamp, 0)

	return data, nil
}

// parseDataToPosition parses a data into position.
func parseDataToPosition(data model.Data) model.Position {
	return model.Position{
		Coordinate: model.Coordinate{
			Latitude:  data.Latitude,
			Longitude: data.Longitude,
		},
		Timestamp: data.Timestamp,
	}
}
