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

type POSStats struct {
	BasicRatio map[string]float64
	ExtendedRatio map[string]map[string]float64
	TokensByPOS map[string][]Token
	UniqueTokensByPOS map[string]map[string]int
}

type ParserResult struct {
	Tokens []Token
	HKKRatio map[string]float64
	POSStats POSStats
}
