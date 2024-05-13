package harvesine

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go-ride-fare-estimation/internal/model"
)

type TestSuite struct {
	suite.Suite
}

type testCase struct {
	name     string
	objs     []model.Coordinate
	distance float64
}

func (suite *TestSuite) TestDistance() {
	testCases := []testCase{
		{
			name: "CalculateDistance",
			objs: []model.Coordinate{
				{
					Latitude:  37.966660,
					Longitude: 23.728308,
				},
				{
					Latitude:  37.966627,
					Longitude: 23.728263,
				},
			},
			distance: 0.005387608950152276,
		},
	}

	for _, tc := range testCases {
		suiteT := suite.T()

		suite.Run(tc.name, func() {
			km, _ := Distance(tc.objs[0], tc.objs[1])

			suite.Equal(tc.distance, km)

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
