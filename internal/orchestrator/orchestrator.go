package orchestrator

import (
	"errors"
	"go-ride-fare-estimation/internal/processor"
	"log"
	"time"
)

// Orchestrator defines the minimum contract our
// orchestrator must satisfy.
type Orchestrator interface {
	Run() error
}

// orchestrator holds the structure of our orch implementation.
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
