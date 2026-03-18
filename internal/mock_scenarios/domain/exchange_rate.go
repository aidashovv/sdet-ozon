package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type RateID string

type ExchangeRate struct {
	RateID    RateID
	NumCode   Code
	CharCode  Code
	Nominal   Nominal
	ValueName ValueName
	Value     Value
	VunitRate VunitRate
}

func NewExchangeRate(rateId, numCode, charCode, valueName, value string, nominal int) (*ExchangeRate, error) {
	checkedNumCode, err := NewCode(numCode)
	if err != nil {
		return nil, err
	}

	checkedCharCode, err := NewCode(charCode)
	if err != nil {
		return nil, err
	}

	checkedNominal, err := NewNominal(nominal)
	if err != nil {
		return nil, err
	}

	checkedValueName, err := NewValueName(valueName)
	if err != nil {
		return nil, err
	}

	checkedValue, err := NewValue(value)
	if err != nil {
		return nil, err
	}

	er := &ExchangeRate{
		RateID:    RateID(rateId),
		NumCode:   checkedNumCode,
		CharCode:  checkedCharCode,
		Nominal:   checkedNominal,
		ValueName: checkedValueName,
		Value:     checkedValue,
	}

	if err := er.updateVunitRate(); err != nil {
		return nil, fmt.Errorf("failed to count vunit rate: %w", err)
	}

	return er, nil
}

func NewExchangeRateFromDB(
	rateId, numCode, charCode, valueName, value string,
	nominal int,
) *ExchangeRate {
	rate, _ := NewExchangeRate(rateId, numCode, charCode, valueName, value, nominal)
	return rate
}

func (er *ExchangeRate) updateVunitRate() error {
	valStr := strings.Replace(er.Value.Amount, ",", ".", 1)
	valFloat, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return fmt.Errorf("invalid parse to float: %w", err)
	}

	res := valFloat / float64(er.Nominal.Amount)
	er.VunitRate.Amount = strings.Replace(fmt.Sprintf("%g", res), ".", ",", 1)

	return nil
}
