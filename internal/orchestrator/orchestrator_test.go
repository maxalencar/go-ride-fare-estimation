package orchestrator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

func BenchmarkRun(b *testing.B) {
	orchestrator, err := NewOrcherstrator("../../test/testdata/paths.csv", "tmpfile.csv")
	if err != nil {
		b.Fail()
	}

	for i := 0; i < b.N; i++ {
		orchestrator.Run()
	}

	if err := os.Remove("tmpfile.csv"); err != nil {
		b.Fail()
	}
}

type TestSuite struct {
	suite.Suite
	orchestrator   Orchestrator
	filePath       string
	resultFilePath string
}

// It runs before each test
func (suite *TestSuite) SetupTest() {
	var err error

	suite.filePath = "../../test/testdata/paths.csv"
	suite.resultFilePath = "tmpfile.csv"
	suite.orchestrator, err = NewOrcherstrator(suite.filePath, suite.resultFilePath)
	if err != nil {
		suite.Fail(err.Error())
	}
}

func (suite *TestSuite) TestRun() {
	suite.orchestrator.Run()
	suite.FileExists(suite.resultFilePath)

	if err := os.Remove(suite.resultFilePath); err != nil {
		suite.Fail(err.Error())
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
