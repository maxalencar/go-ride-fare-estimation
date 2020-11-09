package main

import (
	"flag"
	"go-ride-fare-estimation/internal/orchestrator"
	"go-ride-fare-estimation/internal/processor"
	"log"
)

func main() {
	var fp, rfp string

	flag.StringVar(&fp, "fp", "test/testdata/paths.csv", "file path of the rides positions.")
	flag.StringVar(&rfp, "rfp", "output/result.csv", "file path of the fare estimation results.")
	flag.Parse()

	oc, err := orchestrator.NewOrcherstrator(fp, rfp, processor.NewProcessor())
	if err != nil {
		log.Fatalln(err)
	}

	if err = oc.Run(); err != nil {
		log.Fatalln(err)
	}
}
