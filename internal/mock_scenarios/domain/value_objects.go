package domain

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var (
	ErrCodeLen          = errors.New("num code shoulde be 3")
	ErrNegativeAmount   = errors.New("nominal should be posititve")
	ErrValueNameIsEmpty = errors.New("value name is empty")
	ErrValueIsEmpty     = errors.New("value is empty")
)

type Code struct {
	Code string
}

func NewCode(code string) (Code, error) {
	if utf8.RuneCountInString(code) != 3 {
		return Code{}, ErrCodeLen
	}

	return Code{
		Code: strings.ToUpper(code),
	}, nil
}

type Nominal struct {
	Amount int
}

func NewNominal(amount int) (Nominal, error) {
	if amount < 0 {
		return Nominal{}, ErrNegativeAmount
	}

	return Nominal{
		Amount: amount,
	}, nil
}

type ValueName struct {
	Name string
}

func NewValueName(name string) (ValueName, error) {
	if len(name) == 0 {
		return ValueName{}, ErrValueNameIsEmpty
	}

	return ValueName{
		Name: name,
	}, nil
}

type Value struct {
	Amount string
}

func NewValue(amount string) (Value, error) {
	if len(amount) == 0 {
		return Value{}, ErrValueIsEmpty
	}

	return Value{
		Amount: amount,
	}, nil
}

type VunitRate struct {
	Amount string
}

func NewVunitRate(amount string) (VunitRate, error) {
	if len(amount) == 0 {
		return VunitRate{}, ErrValueIsEmpty
	}

	return VunitRate{
		Amount: amount,
	}, nil
}
