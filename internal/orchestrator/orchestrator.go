package orchestrator

import (
	"errors"
	"go-ride-fare-estimation/internal/processor"
	"log"
	"sync"
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
	processor      processor.Processor
}

// NewOrcherstrator creates a new Orchestrator using given file paths.
func NewOrcherstrator(fp, rfp string) (Orchestrator, error) {
	if fp == "" {
		return nil, errors.New("file path must be provided")
	}

	if rfp == "" {
		return nil, errors.New("result file path must be provided")
	}

	return &orchestrator{
		filePath:       fp,
		resultFilePath: rfp,
		processor:      processor.NewProcessor(),
	}, nil
}

// newOrcherstratorTest creates a new Orchestrator used for unit testing purposes using given file paths and a processor instance enabling mocking capability.
func newOrcherstratorTest(fp, rfp string, processor processor.Processor) (Orchestrator, error) {
	return &orchestrator{
		filePath:       fp,
		resultFilePath: rfp,
		processor:      processor,
	}, nil
}

// Run using concurrency pipeline pattern.
// 1. Read the file
// 2. Collect all data of a ride and start processing
// 3. Create valid segments
// 4. Calculate the fare estimate
// 5. Create a result text file
func (o orchestrator) Run() error {
	var wg sync.WaitGroup

	start := time.Now()

	dChan := o.processor.Read(o.filePath, &wg)
	rpChan := o.processor.Process(dChan, &wg)
	rsChan := o.processor.CreateSegments(rpChan, &wg)
	rfChan := o.processor.CalculateFare(rsChan, &wg)
	o.processor.WriteResult(rfChan, o.resultFilePath, &wg)

	wg.Wait()
	log.Printf("It took %s to finish", time.Since(start))

	return nil
}
