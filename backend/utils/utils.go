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

func KanaToPolivanov(s string) (formatted string) {
	formatted = kana.KanaToRomaji(s)
	formatted = ReplaceAll(
		formatted,
		[]string{"ja", "ju", "jo"},
		[]string{"дзя", "дзю", "дзё"},
	)
	formatted = ReplaceAll(
		formatted,
		[]string{"a", "i", "u", "e", "o", "k", "s", "j", "t", "ch", "n", "h", "m", "r", "w", "g", "z", "d", "b", "p", "y", "f"},
		[]string{"а", "и", "у", "э", "о", "к", "с", "дз", "т", "т", "н", "х", "м", "р", "в", "г", "дз","д", "б", "п", "й", "ф"},
	)
	formatted = ReplaceAll(
		formatted,
		[]string{"ннн", "нна", "нни", "нну", "ннэ", "нно", "ннй", "нй", "нн", "нън"},
		[]string{"нън", "нъа", "нъи", "нъу", "нъэ", "нъо", "нъй", "нъй", "н", "нн"},
	)
	formatted = ReplaceAll(
		formatted,
		[]string{"йа", "йу", "йо", "тс", "сх", "ху", "�"},
		[]string{"я", "ю", "ё", "ц", "с", "фу", ""},
	)
	formatted = ReplaceAll(
		formatted,
		[]string{"нб", "нп", "нм"},
		[]string{"мб", "мп", "мм"},
	)
	formatted = ReplaceAll(
		formatted,
		[]string{"оу", "ёу"},
		[]string{"о:", "ё:"},
	)

	return
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
