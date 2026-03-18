package application

import (
	"context"
	"sdet-ozon/internal/mock_scenarios/domain"
)

type RateService struct {
	repo domain.MockRepository
}

func NewRateService(repo domain.MockRepository) *RateService {
	return &RateService{
		repo: repo,
	}
}

func (s *RateService) GetExchangeRate(ctx context.Context, testID string) (*domain.MockScenario, error) {
	scenario, err := s.repo.GetByTestID(ctx, testID)
	if err != nil {
		return nil, err
	}

	return scenario, nil
}
