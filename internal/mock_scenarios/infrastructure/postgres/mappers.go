package postgres

import "sdet-ozon/internal/mock_scenarios/domain"

func ToDomain(scenarioDTO MockScenarioDTO, rateDTO ExchangeRateDTO) *domain.MockScenario {
	rateDomain := *domain.NewExchangeRateFromDB(
		rateDTO.RateID,
		rateDTO.NumCode,
		rateDTO.CharCode,
		rateDTO.ValueName,
		rateDTO.Value,
		rateDTO.Nominal,
	)

	scenarioDomain := *domain.NewMockScenarioFromDB(
		scenarioDTO.TestID,
		scenarioDTO.Version,
		scenarioDTO.StatusCode,
		&rateDomain,
	)

	return &scenarioDomain
}

func ToDBModel(scenario *domain.MockScenario) (MockScenarioDTO, ExchangeRateDTO) {
	scenarioDTO := NewMockScenarioDTO(
		string(scenario.TestID),
		int(scenario.Version),
		int(scenario.StatusCode),
	)

	var rateDTO ExchangeRateDTO
	if scenario.Data != nil {
		rateDTO = ExchangeRateDTO{
			RateID:    string(scenario.Data.RateID),
			NumCode:   string(scenario.Data.NumCode.Code),
			CharCode:  string(scenario.Data.CharCode.Code),
			Nominal:   scenario.Data.Nominal.Amount,
			ValueName: scenario.Data.ValueName.Name,
			Value:     scenario.Data.Value.Amount,
			VunitRate: scenario.Data.VunitRate.Amount,
		}
	}

	return scenarioDTO, rateDTO
}
