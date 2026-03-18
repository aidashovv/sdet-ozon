package presentation

import (
	"encoding/xml"
)

type SetupMockDTO struct {
	TestID     string                `json:"test_id"`
	StatusCode int                   `json:"status_code"`
	Rate       CreateExchangeRateDTO `json:"rate"`
}

type CreateExchangeRateDTO struct {
	RateID    string `json:"rate_id" xml:"ID,attr"`
	NumCode   string `json:"num_code" xml:"NumCode"`
	CharCode  string `json:"char_code" xml:"CharCode"`
	Nominal   int    `json:"nominal" xml:"Nominal"`
	ValueName string `json:"value_name" xml:"Name"`
	Value     string `json:"value" xml:"Value"`
}

type ResponseExchangeRateDTO struct {
	RateID    string `json:"rate_id" xml:"ID,attr"`
	NumCode   string `json:"num_code" xml:"NumCode"`
	CharCode  string `json:"char_code" xml:"CharCode"`
	Nominal   int    `json:"nominal" xml:"Nominal"`
	ValueName string `json:"value_name" xml:"Name"`
	Value     string `json:"value" xml:"Value"`
	VunitRate string `json:"vunit_rate" xml:"VunitRate"`
}

type ValCursXML struct {
	XMLName xml.Name                  `xml:"ValCurs"`
	Date    string                    `xml:"Date,attr"`
	Name    string                    `xml:"name,attr"`
	Valutes []ResponseExchangeRateDTO `xml:"Valute"`
}

func (v ValCursXML) ToXML() []byte {
	output, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	return append([]byte(xml.Header), output...)
}
