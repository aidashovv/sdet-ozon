package postgres

type MockScenarioDTO struct {
	TestID     string `db:"test_id"`
	Version    int    `db:"version"`
	StatusCode int    `db:"status_code"`
}

func NewMockScenarioDTO(id string, version, statusCode int) MockScenarioDTO {
	return MockScenarioDTO{
		TestID:     id,
		Version:    version,
		StatusCode: statusCode,
	}
}

type ExchangeRateDTO struct {
	RateID    string `db:"rate_id"`
	NumCode   string `db:"num_code"`
	CharCode  string `db:"char_code"`
	Nominal   int    `db:"nominal"`
	ValueName string `db:"value_name"`
	Value     string `db:"value"`
	VunitRate string `db:"vunit_rate"`
}

func NewExchangeRateDTO(
	rateId, numCode, charCode, valueName, value, vunitRate string,
	nominal int,
) ExchangeRateDTO {
	return ExchangeRateDTO{
		RateID:    rateId,
		NumCode:   numCode,
		CharCode:  charCode,
		Nominal:   nominal,
		ValueName: valueName,
		Value:     value,
		VunitRate: vunitRate,
	}
}
