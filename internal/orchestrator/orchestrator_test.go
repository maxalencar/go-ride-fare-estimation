package orchestrator

import (
	"go-ride-fare-estimation/internal/processor"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	orchestrator Orchestrator
	fp           string
	rfp          string
}

// It runs before each test
func (suite *TestSuite) SetupTest() {
	var err error

	suite.fp = "../../test/testdata/paths-small.csv"
	suite.rfp = "tmpfile.csv"
	suite.orchestrator, err = NewOrcherstrator(suite.fp, suite.rfp, processor.NewProcessor())
	if err != nil {
		suite.Fail(err.Error())
	}
}

func (suite *TestSuite) TestRun() {
	suite.orchestrator.Run()
	suite.FileExists(suite.rfp)

	if err := os.Remove(suite.rfp); err != nil {
		suite.Fail(err.Error())
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
