package exporter

import (
	"os"
	"strconv"
	"strings"
	"github.com/programmer8760/japanese-parser/backend/types"
	"github.com/programmer8760/japanese-parser/backend/utils"
)

func ExportTxt(parserResult types.ParserResult, path string) error {
	data := []byte("Оригинал\n")
	reading := []byte("\n\nФуригана\n")
	romaji := []byte("\n\nРомадзи\n")
	kiriji := []byte("\n\nКиридзи\n")
	
	for _, token := range parserResult.Tokens {
		data = append(data, []byte(token.Surface)...)
		if token.Reading != "*" {
			reading = append(reading, []byte(token.Reading + " ")...)
		} else {
			reading = append(reading, []byte(token.Surface + " ")...)
		}
		if token.Romaji != "*" {
			romaji = append(romaji, []byte(token.Romaji + " ")...)
		} else {
			romaji = append(romaji, []byte(token.Surface + " ")...)
		}
		if token.Polivanov != "*" {
			kiriji = append(kiriji, []byte(token.Polivanov + " ")...)
		} else {
			kiriji = append(kiriji, []byte(token.Surface + " ")...)
		}
	}
	data = append(data, reading...)
	data = append(data, romaji...)
	data = append(data, kiriji...)

	hkkRatio := []byte("\n\nСоотношение символов\nХирагана: ")
	hiragana := strconv.FormatFloat(parserResult.HKKRatio["hiragana"], 'f', 2, 64)
	hkkRatio = append(hkkRatio, []byte(hiragana)...)
	hkkRatio = append(hkkRatio, []byte("%\nКатакана: ")...)
	katakana := strconv.FormatFloat(parserResult.HKKRatio["katakana"], 'f', 2, 64)
	hkkRatio = append(hkkRatio, []byte(katakana)...)
	hkkRatio = append(hkkRatio, []byte("%\nКандзи: ")...)
	kanji := strconv.FormatFloat(parserResult.HKKRatio["kanji"], 'f', 2, 64)
	hkkRatio = append(hkkRatio, []byte(kanji)...)
	hkkRatio = append(hkkRatio, []byte("%")...)

	data = append(data, hkkRatio...)

	posRatio := []byte("\n\nСоотношение частей речи\n")
	for pos, value := range parserResult.POSStats.BasicRatio {
		posRatio = append(posRatio, []byte(pos + " ")...)
		posRatio = append(posRatio, []byte(strconv.FormatFloat(value, 'f', 2, 64) + "%: ")...)
		for subPOS, subValue := range parserResult.POSStats.ExtendedRatio[pos] {
			posRatio = append(posRatio, []byte(subPOS + " ")...)
			posRatio = append(posRatio, []byte(strconv.FormatFloat(subValue, 'f', 2, 64) + "%, ")...)
		}
		posRatio = append(posRatio, []byte("\n")...)
	}

	data = append(data, posRatio...)

	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ExportCsv(parserResult types.ParserResult, path string) error {
	data := []byte("Токен;Чтение;Ромадзи;Киридзи;Часть речи;Подкласс;Перевод\n")
	for _, tokens := range utils.GetUniqueTokens(parserResult.POSStats) {
		for _, token := range tokens {
			translations := "Значение не найдено"
			if token.Translations != nil {
				translations = strings.Join(token.Translations[0].Translations, " // ")
			}
			var entry string
			if token.BaseForm != token.Surface && token.BaseForm != "*" {
				entry = strings.Join([]string{
					token.BaseForm,
					token.BaseFormReading,
					token.BaseFormRomaji,
					token.BaseFormPolivanov,
					token.POS[0],
					token.POS[1],
					translations,
				}, ";")
			} else {
				entry = strings.Join([]string{
					token.Surface,
					token.Reading,
					token.Romaji,
					token.Polivanov,
					token.POS[0],
					token.POS[1],
					translations,
				}, ";")
			}

			data = append(data, []byte(entry + "\n")...)
		}
	}

	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
