package analyzer

import (
	"github.com/programmer8760/japanese-parser/backend/types"

	"github.com/gojp/kana"

	"unicode/utf8"
)

func GetHKKRatio(tokens []types.Token) map[string]float64 {
	total, hiragana, katakana, kanji := 0, 0, 0, 0
	for _, t := range tokens {
		if kana.IsKana(t.Surface) || kana.IsKanji(t.Surface) {
			total += utf8.RuneCountInString(t.Surface)
			switch {
			case kana.IsHiragana(t.Surface):
				hiragana += utf8.RuneCountInString(t.Surface)
			case kana.IsKatakana(t.Surface):
				katakana += utf8.RuneCountInString(t.Surface)
			case kana.IsKanji(t.Surface):
				kanji += utf8.RuneCountInString(t.Surface)
			}
		} else {
			for _, r := range t.Surface {
				switch {
				case kana.IsHiragana(string(r)):
					total += 1
					hiragana += 1
				case kana.IsKatakana(string(r)):
					total += 1
					katakana += 1
				case kana.IsKanji(string(r)):
					total += 1
					kanji += 1
				}
			}
		}
	}

	return map[string]float64{
		"hiragana": float64(hiragana)*100/float64(total),
		"katakana": float64(katakana)*100/float64(total),
		"kanji": float64(kanji)*100/float64(total),
	}
}

func GetPOSStats(tokens []types.Token) types.POSStats {
	basic := make(map[string]float64)
	extended := make(map[string]map[string]float64)
	tokensByPOS := make(map[string][]types.Token)
	uniqueTokensByPOS := make(map[string]map[string]int)
	total := 0
	for _, t := range tokens {
		if t.POS[0] == "記号" { //skip whitespaces, dots and other symbols
			continue
		}

		basic[t.POS[0]] += 1

		if _, ok := extended[t.POS[0]]; !ok {
			extended[t.POS[0]] = make(map[string]float64)
		}
		extended[t.POS[0]][t.POS[1]] += 1

		tokensByPOS[t.POS[0]] = append(tokensByPOS[t.POS[0]], t)

		if _, ok := uniqueTokensByPOS[t.POS[0]]; !ok {
			uniqueTokensByPOS[t.POS[0]] = make(map[string]int)
		}
		if _, exists := uniqueTokensByPOS[t.POS[0]][t.BaseForm]; !exists {
			uniqueTokensByPOS[t.POS[0]][t.BaseForm] = len(tokensByPOS[t.POS[0]]) - 1
		}

		total += 1
	}
	for posName, subPos := range extended {
		for key, value := range subPos {
			subPos[key] = value*100/basic[posName]
		}
		basic[posName] = basic[posName]*100/float64(total)
	}
	return types.POSStats{
		BasicRatio: basic,
		ExtendedRatio: extended,
		TokensByPOS: tokensByPOS,
		UniqueTokensByPOS: uniqueTokensByPOS,
	}
}
