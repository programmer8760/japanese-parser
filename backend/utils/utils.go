package utils

import (
	"github.com/gojp/kana"
	"strings"
)

func KatakanaToHiragana(s string) string {
    runes := []rune(s)
    for i, r := range runes {
        if r >= 0x30A1 && r <= 0x30F6 {
            runes[i] = r - 0x60
        }
    }
    return string(runes)
}

func ReplaceAll(haystack string, needles []string, replacements []string) (replaced string) {
	replaced = haystack
	for i := range needles {
		replaced = strings.Replace(replaced, needles[i], replacements[i], -1)
	}
	return replaced
}

func KanaToRomaji(s string) string {
	return ReplaceAll(kana.KanaToRomaji(s), []string{"nn", "�"}, []string{"n'", ""})
}
