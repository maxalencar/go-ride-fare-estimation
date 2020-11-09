package e2e

import (
	"encoding/csv"
	"go-ride-fare-estimation/internal/orchestrator"
	"go-ride-fare-estimation/internal/processor"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

func BenchmarkRun(b *testing.B) {
	orchestrator, err := orchestrator.NewOrcherstratorTest("../testdata/paths.csv", "tmpfile.csv", processor.NewProcessor(), sync.WaitGroup{})
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
	orchestrator orchestrator.Orchestrator
	fp           string
	rfp          string
}

// It runs before each test
func (suite *TestSuite) SetupTest() {
	var err error

	suite.fp = "../testdata/paths.csv"
	suite.rfp = "tmpfile.csv"
	suite.orchestrator, err = orchestrator.NewOrcherstratorTest(suite.fp, suite.rfp, processor.NewProcessor(), sync.WaitGroup{})
	if err != nil {
		suite.Fail(err.Error())
	}
}

func (suite *TestSuite) TestRun() {
	suite.orchestrator.Run()
	suite.FileExists(suite.rfp)

	f, err := os.Open(suite.rfp)
	if err != nil {
		os.Remove(suite.rfp)
		suite.Fail(err.Error())
	}

	l, err := csv.NewReader(f).ReadAll()
	if err != nil {
		os.Remove(suite.rfp)
		suite.Fail(err.Error())
	}
	f.Close()

	suite.Equal(9, len(l))
	suite.Equal("1", l[0][0])
	suite.Equal("11.3391406912", l[0][1])
	suite.Equal("2", l[1][0])
	suite.Equal("13.0992278125", l[1][1])
	suite.Equal("3", l[2][0])
	suite.Equal("33.8442543273", l[2][1])
	suite.Equal("4", l[3][0])
	suite.Equal("3.4700000000", l[3][1])
	suite.Equal("5", l[4][0])
	suite.Equal("22.7758294424", l[4][1])
	suite.Equal("6", l[5][0])
	suite.Equal("9.4142865094", l[5][1])
	suite.Equal("7", l[6][0])
	suite.Equal("30.0108781970", l[6][1])
	suite.Equal("8", l[7][0])
	suite.Equal("9.2085952581", l[7][1])
	suite.Equal("9", l[8][0])
	suite.Equal("6.3471241656", l[8][1])

	if err := os.Remove(suite.rfp); err != nil {
		suite.Fail(err.Error())
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
