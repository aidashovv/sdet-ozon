package application

import (
	"context"
	"errors"
	"fmt"
	"sdet-ozon/internal/mock_scenarios/domain"
	"sdet-ozon/internal/pkg/myerr"
)

type SetupService struct {
	repo domain.MockRepository
}

func NewSetupService(repo domain.MockRepository) *SetupService {
	return &SetupService{
		repo: repo,
	}
}

func (s *SetupService) RegisterScenario(ctx context.Context, scenario *domain.MockScenario) error {
	existing, err := s.repo.GetByTestID(ctx, string(scenario.TestID))

	if err != nil {
		if errors.Is(err, myerr.ErrScenarioNotFound) {
			return s.repo.Add(ctx, scenario)
		}
		return fmt.Errorf("check existing scenario: %w", err)
	}

	existing.UpdateFrom(scenario)

	if err := s.repo.Update(ctx, existing); err != nil {
		return fmt.Errorf("update scenario: %w", err)
	}

	return nil
}

func (s *SetupService) DeleteScenario(ctx context.Context, testID string) error {
	return s.repo.Delete(ctx, testID)
}
