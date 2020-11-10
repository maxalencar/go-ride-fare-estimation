package processor

import (
	"go-ride-fare-estimation/internal/model"
	"sync"

	"github.com/stretchr/testify/mock"
)

// Mock - mock
type Mock struct {
	mock.Mock
}

// CalculateFare - mock
func (m *Mock) CalculateFare(in <-chan model.Ride, wg *sync.WaitGroup) <-chan model.Ride {
	args := m.Called(in, wg)

	return args.Get(0).(chan model.Ride)
}

// CreateSegments - mock
func (m *Mock) CreateSegments(in <-chan model.Ride, wg *sync.WaitGroup) <-chan model.Ride {
	args := m.Called(in, wg)

	return args.Get(0).(chan model.Ride)
}

// Process - mock
func (m *Mock) Process(in <-chan model.Data, wg *sync.WaitGroup) <-chan model.Ride {
	args := m.Called(in, wg)

	return args.Get(0).(chan model.Ride)
}

// Read - mock
func (m *Mock) Read(filePath string, wg *sync.WaitGroup) <-chan model.Data {
	args := m.Called(filePath, wg)

	return args.Get(0).(chan model.Data)
}

// WriteResult - mock
func (m *Mock) WriteResult(in <-chan model.Ride, filePath string, wg *sync.WaitGroup) {
	m.Called(in, filePath, wg)
}
