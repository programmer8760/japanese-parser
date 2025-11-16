package analyzer

import (
	"github.com/programmer8760/japanese-parser/backend/types"

	"github.com/gojp/kana"

	"unicode/utf8"
)

func GetHKKRatio(tokens []types.Token) []int {
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

	return []int{
		hiragana*100/total,
		katakana*100/total,
		kanji*100/total,
	}
}
