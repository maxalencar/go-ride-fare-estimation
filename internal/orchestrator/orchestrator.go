package orchestrator

import (
	"errors"
	"go-ride-fare-estimation/internal/processor"
	"log"
	"time"
)

// Orchestrator defines the minimum contract our orchestrator must satisfy.
type Orchestrator interface {
	Run() error
}

// orchestrator holds the structure of our orchestrator implementation.
type orchestrator struct {
	filePath       string
	resultFilePath string
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
	}, nil
}

// Run using concurrency pipeline pattern.
// 1. Read the file
// 2. Collect all data of a ride and start processing
// 3. Create valid segments
// 4. Calculate the fare estimate
// 5. Create a result text file
func (o orchestrator) Run() error {
	start := time.Now()

	proc := processor.NewProcessor()

	c := proc.Read(o.filePath)
	r := proc.Process(c)
	s := proc.CreateSegments(r)
	f := proc.CalculateFare(s)

	err := proc.WriteResult(f, o.resultFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("It took %s", time.Since(start))

	return nil
}
