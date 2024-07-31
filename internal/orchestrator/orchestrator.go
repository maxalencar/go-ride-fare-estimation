package orchestrator

import (
	"errors"
	"log"
	"sync"
	"time"

	"go-ride-fare-estimation/internal/processor"
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
	wg             sync.WaitGroup
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

// Run using concurrency pipeline pattern.
// 1. Read the file
// 2. Collect all data of a ride and start processing
// 3. Create valid segments
// 4. Calculate the fare estimate
// 5. Create a result text file
func (o *orchestrator) Run() error {
	start := time.Now()

	dChan := o.processor.Read(o.filePath, &o.wg)
	rpChan := o.processor.Process(dChan, &o.wg)
	rsChan := o.processor.CreateSegments(rpChan, &o.wg)
	rfChan := o.processor.CalculateFare(rsChan, &o.wg)
	o.processor.WriteResult(rfChan, o.resultFilePath, &o.wg)

	o.wg.Wait()
	log.Printf("It took %s to finish", time.Since(start))

	return nil
}
