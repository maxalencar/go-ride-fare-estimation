package segment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"go-ride-fare-estimation/internal/model"
)

type TestSuite struct {
	suite.Suite
}

type testCase struct {
	name  string
	objs  []model.Position
	count int
}

func (suite *TestSuite) TestCreate() {
	testCases := []testCase{
		{
			name: "CreateOneValidSegment",
			objs: []model.Position{
				{
					Coordinate: model.Coordinate{
						Latitude:  37.966660,
						Longitude: 23.728308,
					},
					Timestamp: time.Unix(1405594957, 0),
				},
				{
					Coordinate: model.Coordinate{
						Latitude:  37.966627,
						Longitude: 23.728263,
					},
					Timestamp: time.Unix(1405594966, 0),
				},
			},
			count: 1,
		},
		{
			name: "CreateTwoValidSegmentsWithOneInvalidPosition",
			objs: []model.Position{
				{
					Coordinate: model.Coordinate{
						Latitude:  37.966660,
						Longitude: 23.728308,
					},
					Timestamp: time.Unix(1405594957, 0),
				},
				{
					Coordinate: model.Coordinate{
						Latitude:  37.966627,
						Longitude: 23.728263,
					},
					Timestamp: time.Unix(1405594966, 0),
				},
				{
					Coordinate: model.Coordinate{
						Latitude:  37.967627,
						Longitude: 30.729563,
					},
					Timestamp: time.Unix(1405594976, 0),
				},
				{
					Coordinate: model.Coordinate{
						Latitude:  37.966647,
						Longitude: 23.728763,
					},
					Timestamp: time.Unix(1405594986, 0),
				},
			},
			count: 2,
		},
		{
			name: "CreateOneValidSegmentWithLastInvalidPosition",
			objs: []model.Position{
				{
					Coordinate: model.Coordinate{
						Latitude:  37.966660,
						Longitude: 23.728308,
					},
					Timestamp: time.Unix(1405594957, 0),
				},
				{
					Coordinate: model.Coordinate{
						Latitude:  37.966627,
						Longitude: 23.728263,
					},
					Timestamp: time.Unix(1405594966, 0),
				},
				{
					Coordinate: model.Coordinate{
						Latitude:  37.967627,
						Longitude: 30.729563,
					},
					Timestamp: time.Unix(1405594976, 0),
				},
			},
			count: 1,
		},
	}

	for _, tc := range testCases {
		suiteT := suite.T()

		suite.Run(tc.name, func() {
			segments := Create(tc.objs)

			suite.Equal(tc.count, len(segments))

			// We should get a different *testing.T for subTests, so that
			// go test recognises them as proper subtests for output formatting
			// and running individuals subtests
			subTest := suite.T()
			suite.NotEqual(subTest, suiteT)
		})
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
