package domain

import "context"

type MockRepository interface {
	Add(ctx context.Context, scenario *MockScenario) error
	GetByTestID(ctx context.Context, id string) (*MockScenario, error)
	Update(ctx context.Context, scenario *MockScenario) error
	Delete(ctx context.Context, id string) error
}
