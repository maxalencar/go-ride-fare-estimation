package fare

import (
	"go-ride-fare-estimation/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

type testCase struct {
	name     string
	objs     []model.Segment
	expected interface{}
}

func (suite *TestSuite) TestCalculate() {
	testCases := []testCase{
		{
			name: "CalculateIdleAndLessThanMinimum",
			objs: []model.Segment{
				model.Segment{
					Position1: model.Position{
						Coordinate: model.Coordinate{
							Latitude:  37.966660,
							Longitude: 23.728308,
						},
						Timestamp: time.Unix(1405594957, 0),
					},
					Position2: model.Position{
						Coordinate: model.Coordinate{
							Latitude:  37.966627,
							Longitude: 23.728263,
						},
						Timestamp: time.Unix(1405594966, 0),
					},
					Distance: 0.005387608950152276,
					Duration: time.Duration(9 * time.Second),
					Speed:    2.1550435800609105,
				},
			},
			expected: 3.47,
		},
		{
			name: "CalculateMovingNormalFare",
			objs: []model.Segment{
				model.Segment{
					Position1: model.Position{
						Timestamp: time.Unix(1405594957, 0),
					},
					Position2: model.Position{
						Timestamp: time.Unix(1405594966, 0),
					},
					Distance: 3.8,
					Speed:    80,
				},
				model.Segment{
					Position1: model.Position{
						Timestamp: time.Unix(1405594957, 0),
					},
					Position2: model.Position{
						Timestamp: time.Unix(1405594966, 0),
					},
					Distance: 2.5,
					Speed:    50,
				},
			},
			expected: 5.962,
		},
		{
			name: "CalculateMovingExtraFare",
			objs: []model.Segment{
				model.Segment{
					Position1: model.Position{
						Timestamp: time.Unix(1604794500, 0),
					},
					Position2: model.Position{
						Timestamp: time.Unix(1604794500, 0),
					},
					Distance: 3.8,
					Speed:    80,
				},
				model.Segment{
					Position1: model.Position{
						Timestamp: time.Unix(1604794500, 0),
					},
					Position2: model.Position{
						Timestamp: time.Unix(1604794500, 0),
					},
					Distance: 2.5,
					Speed:    50,
				},
			},
			expected: 9.489999999999998,
		},
		{
			name: "CalculateMovingNormalAndExtraFare",
			objs: []model.Segment{
				model.Segment{
					Position1: model.Position{
						Timestamp: time.Unix(1405594957, 0),
					},
					Position2: model.Position{
						Timestamp: time.Unix(1405594957, 0),
					},
					Distance: 3.8,
					Speed:    80,
				},
				model.Segment{
					Position1: model.Position{
						Timestamp: time.Unix(1604794500, 0),
					},
					Position2: model.Position{
						Timestamp: time.Unix(1604794500, 0),
					},
					Distance: 2.5,
					Speed:    50,
				},
			},
			expected: 7.362,
		},
	}

	for _, tc := range testCases {
		suiteT := suite.T()

		suite.Run(tc.name, func() {
			fare := Calculate(tc.objs)

			suite.Equal(tc.expected, fare)

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
