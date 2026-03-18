package domain

import "errors"

var (
	ErrRequiredData = errors.New("exchange rate data is required for 200 status")
)

type TestID string
type Version int
type StatusCode int

type MockScenario struct {
	TestID     TestID
	Version    Version
	StatusCode StatusCode
	Data       *ExchangeRate
}

func NewMockScenario(id string, statusCode int, data *ExchangeRate) (*MockScenario, error) {
	if statusCode == 200 && data == nil {
		return nil, ErrRequiredData
	}

	return &MockScenario{
		TestID:     TestID(id),
		Version:    Version(1),
		StatusCode: StatusCode(statusCode),
		Data:       data,
	}, nil
}

func NewMockScenarioFromDB(id string, version, statusCode int, data *ExchangeRate) *MockScenario {
	scneraio, _ := NewMockScenario(id, statusCode, data)
	scneraio.Version = Version(version)

	return scneraio
}

func (s *MockScenario) UpdateFrom(other *MockScenario) {
	s.StatusCode = other.StatusCode

	if other.Data == nil {
		s.Data = nil
		return
	}

	if s.Data == nil {
		s.Data = other.Data
	} else {
		*s.Data = *other.Data
	}
}
