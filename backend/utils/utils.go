package utils

func KatakanaToHiragana(s string) string {
    runes := []rune(s)
    for i, r := range runes {
        if r >= 0x30A1 && r <= 0x30F6 {
            runes[i] = r - 0x60
        }
    }
    return string(runes)
}
