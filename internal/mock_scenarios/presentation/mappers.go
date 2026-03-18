package presentation

import (
	"sdet-ozon/internal/mock_scenarios/domain"
	"strings"
	"time"
)

func ToDomain(dto SetupMockDTO) (*domain.MockScenario, error) {
	var rate *domain.ExchangeRate

	if dto.StatusCode == 200 {
		rate, _ = domain.NewExchangeRate(
			dto.Rate.RateID,
			dto.Rate.NumCode,
			dto.Rate.CharCode,
			dto.Rate.ValueName,
			dto.Rate.Value,
			dto.Rate.Nominal,
		)
	}

	scenario, err := domain.NewMockScenario(dto.TestID, dto.StatusCode, rate)
	if err != nil {
		return nil, err
	}

	return scenario, nil
}

func ToXMLResponse(scenario *domain.MockScenario) ValCursXML {
	response := ValCursXML{
		Date: time.Now().Format("30/03/2006"),
		Name: "ххх",
	}

	if scenario.Data == nil {
		return response
	}

	formattedValue := strings.ReplaceAll(scenario.Data.Value.Amount, ".", ",")

	vunitRaw := scenario.Data.VunitRate.Amount
	formattedVunit := strings.ReplaceAll(vunitRaw, ".", ",")

	response.Valutes = []ResponseExchangeRateDTO{
		{
			RateID:    string(scenario.Data.RateID),
			NumCode:   scenario.Data.NumCode.Code,
			CharCode:  scenario.Data.CharCode.Code,
			Nominal:   scenario.Data.Nominal.Amount,
			ValueName: scenario.Data.ValueName.Name,
			Value:     formattedValue,
			VunitRate: formattedVunit,
		},
	}

	return response
}
