package orchestrator

import (
	"errors"
	"go-ride-fare-estimation/internal/fare"
	"go-ride-fare-estimation/internal/model"
	"go-ride-fare-estimation/internal/processor"
	"go-ride-fare-estimation/internal/segment"
	"log"
	"time"
)

// Orchestrator defines the minimum contract our
// orchestrator must satisfy.
type Orchestrator interface {
	Run() error
}

// orchestrator holds the structure of our TCP implementation.
type orchestrator struct {
	filePath       string
	resultFilePath string
	csegs          chan *model.Ride
	cfares         chan *model.Ride
	rides          map[int]*model.Ride
}

// NewOrcherstrator creates a new Orchestrator using given file path.
func NewOrcherstrator(fp, rfp string) (Orchestrator, error) {
	if fp == "" {
		return nil, errors.New("a file path must be provided")
	}

	if rfp == "" {
		return nil, errors.New("a result file path must be provided")
	}

	return &orchestrator{
		filePath:       fp,
		resultFilePath: rfp,
		csegs:          make(chan *model.Ride),
		cfares:         make(chan *model.Ride),
	}, nil
}

// broadcaster - it broadcasts messages based on the selected channel
func (o orchestrator) broadcaster() {
	for {
		select {
		case ride := <-o.csegs:
			log.Printf("creating segments for ride %d \n", ride.ID)
			ride.Segments = segment.Create(ride.Positions)
		case ride := <-o.cfares:
			log.Printf("calculating fare for ride %d \n", ride.ID)
			ride.FareEstimate = fare.Calculate(ride.Segments)
		}
	}
}

func (o orchestrator) Run() error {
	start := time.Now()

	proc := processor.NewProcessor()

	c := proc.Read(o.filePath)
	r := proc.Process(c)
	s := proc.CreateSegments(r)
	f := proc.CalculateFare(s)

	proc.WriteResult(f, o.resultFilePath)

	log.Printf("It took %s", time.Since(start))

	return nil
}
