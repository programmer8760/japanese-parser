package types

type Token struct {
	Surface string
	POS []string
	BaseForm string
	InflectionalForm string
	InflectionalType string
	Translations []DictionaryEntry
	Reading string
	Romaji string
	Polivanov string
}

type DictionaryEntry struct {
	Kanji string
	Reading string
	Translations []string
	WordID int
}

type ParserResult struct {
	Tokens []Token
	HKKRatio []int
	POSRatio []int
	UniqueWords []Token
}
