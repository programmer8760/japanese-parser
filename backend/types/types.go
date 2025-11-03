package types

type Token struct {
	Surface string `json:"surface"`
	Pronunciation string `json:"pronuncition"`
	POS []string `json:"pos"`
	BaseForm string `json:"base_form"`
	InflectionalForm string `json:"inflectional_form"`
	InflectionalType string `json:"inflectional_type"`
	Translation string `json:"translation"`
}
