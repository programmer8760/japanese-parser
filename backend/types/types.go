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
	BasicRatio map[string]int
	ExtendedRatio map[string]map[string]int
	TokensByPOS map[string][]Token
	UniqueTokensByPOS map[string]map[string]int
}

type ParserResult struct {
	Tokens []Token
	HKKRatio map[string]int
	POSStats POSStats
}
