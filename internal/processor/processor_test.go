package processor

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"go-ride-fare-estimation/internal/model"
)

type TestSuite struct {
	suite.Suite
	processor Processor
}

// It runs before each test
func (suite *TestSuite) SetupTest() {
	suite.processor = NewProcessor()
}
func (suite *TestSuite) TestRead() {
	var wg sync.WaitGroup

	dChan := suite.processor.Read("../../test/testdata/paths.csv", &wg)

	dataList := []model.Data{}

	for d := range dChan {
		dataList = append(dataList, d)
	}

	wg.Wait()
	suite.Equal(1826, len(dataList))
}

func (suite *TestSuite) TestProcess() {
	var wg sync.WaitGroup

	var dChan = make(chan model.Data, 3)

	rChan := suite.processor.Process(dChan, &wg)

	dChan <- model.Data{RideID: 1, Latitude: 37.966660, Longitude: 23.728308, Timestamp: time.Unix(1405594957, 0)}
	dChan <- model.Data{RideID: 1, Latitude: 37.966627, Longitude: 23.728263, Timestamp: time.Unix(1405594966, 0)}
	dChan <- model.Data{RideID: 2, Latitude: 37.966627, Longitude: 23.728263, Timestamp: time.Unix(1405594966, 0)}
	close(dChan)

	rides := []model.Ride{}

	for r := range rChan {
		fmt.Println(r)
		rides = append(rides, r)
	}

	wg.Wait()
	suite.Equal(2, len(rides))
	suite.Equal(1, rides[0].ID)
	suite.Equal(2, len(rides[0].Positions))
	suite.Equal(2, rides[1].ID)
	suite.Equal(1, len(rides[1].Positions))
}

func (suite *TestSuite) TestCreateSegment() {
	var wg sync.WaitGroup

	var rChan = make(chan model.Ride, 2)

	r2Chan := suite.processor.CreateSegments(rChan, &wg)

	rChan <- model.Ride{ID: 1, Positions: []model.Position{
		{Coordinate: model.Coordinate{Latitude: 37.966660, Longitude: 23.728308}, Timestamp: time.Unix(1405594957, 0)},
		{Coordinate: model.Coordinate{Latitude: 37.966627, Longitude: 23.728263}, Timestamp: time.Unix(1405594966, 0)},
	}}
	rChan <- model.Ride{ID: 2, Positions: []model.Position{
		{Coordinate: model.Coordinate{Latitude: 37.966627, Longitude: 23.728263}, Timestamp: time.Unix(1405594966, 0)},
	}}
	close(rChan)

	rides := []model.Ride{}

	for r := range r2Chan {
		rides = append(rides, r)
	}

	wg.Wait()
	suite.Equal(2, len(rides))
	suite.Equal(1, rides[0].ID)
	suite.Equal(1, len(rides[0].Segments))
	suite.Equal(2, rides[1].ID)
	suite.Equal(0, len(rides[1].Segments))
}

func (suite *TestSuite) TestCalculateFare() {
	var wg sync.WaitGroup

	var rChan = make(chan model.Ride, 2)

	r2Chan := suite.processor.CalculateFare(rChan, &wg)

	rChan <- model.Ride{ID: 1, Segments: []model.Segment{
		{
			Position1: model.Position{
				Timestamp: time.Unix(1405594957, 0),
			},
			Position2: model.Position{
				Timestamp: time.Unix(1405594957, 0),
			},
			Distance: 3.8,
			Speed:    80,
		},
		{
			Position1: model.Position{
				Timestamp: time.Unix(1604794500, 0),
			},
			Position2: model.Position{
				Timestamp: time.Unix(1604794500, 0),
			},
			Distance: 2.5,
			Speed:    50,
		},
	}}
	rChan <- model.Ride{ID: 2}
	close(rChan)

	rides := []model.Ride{}

	for r := range r2Chan {
		rides = append(rides, r)
	}

	wg.Wait()
	suite.Equal(2, len(rides))
	suite.Equal(1, rides[0].ID)
	suite.Equal(7.362, rides[0].FareEstimate)
	suite.Equal(2, rides[1].ID)
	suite.Equal(3.47, rides[1].FareEstimate)
}

func (suite *TestSuite) TestWriteResult() {
	fp := "tmpfile.csv"
	var wg sync.WaitGroup

	var rChan = make(chan model.Ride, 2)

	suite.processor.WriteResult(rChan, fp, &wg)

	rChan <- model.Ride{ID: 1, FareEstimate: 100}
	rChan <- model.Ride{ID: 2, FareEstimate: 200}
	close(rChan)

	wg.Wait()
	suite.FileExists(fp)

	f, err := os.Open(fp)
	if err != nil {
		os.Remove(fp)
		suite.Fail(err.Error())
	}

	l, err := csv.NewReader(f).ReadAll()
	if err != nil {
		os.Remove(fp)
		suite.Fail(err.Error())
	}
	f.Close()

	suite.Equal(2, len(l))
	suite.Equal("1", l[0][0])
	suite.Equal("100.0000000000", l[0][1])
	suite.Equal("2", l[1][0])
	suite.Equal("200.0000000000", l[1][1])

	if err := os.Remove(fp); err != nil {
		suite.Fail(err.Error())
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
