package e2e

import (
	"encoding/csv"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	appPath        string
	filePath       string
	resultFilePath string
}

// It runs before each test
func (suite *TestSuite) SetupTest() {
	suite.appPath = "../../cmd/go-ride-fare-estimation/main.go"
	suite.filePath = "../testdata/paths.csv"
	suite.resultFilePath = "tmpfile.csv"
}

func (suite *TestSuite) TestMain() {
	cmd := exec.Command("go", "run", suite.appPath, "-fp", suite.filePath, "-rfp", suite.resultFilePath)

	if err := cmd.Run(); err != nil {
		os.Remove(suite.resultFilePath)
		suite.Fail(err.Error())
	}

	suite.FileExists(suite.resultFilePath)

	f, err := os.Open(suite.resultFilePath)
	if err != nil {
		os.Remove(suite.resultFilePath)
		suite.Fail(err.Error())
	}

	l, err := csv.NewReader(f).ReadAll()
	if err != nil {
		os.Remove(suite.resultFilePath)
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

	if err := os.Remove(suite.resultFilePath); err != nil {
		suite.Fail(err.Error())
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
