package parser

import (
	"github.com/programmer8760/japanese-parser/backend/types"
	"github.com/programmer8760/japanese-parser/backend/dictionary"

	"github.com/ikawaha/kagome-dict/ipa"
  "github.com/ikawaha/kagome/v2/tokenizer"

	"github.com/gojp/kana"
)

type Parser struct {
	tokenizer *tokenizer.Tokenizer
	dictionary *dictionary.Dictionary
}

func NewParser() (*Parser, error) {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return nil, err
	}
	d, err := dictionary.NewDictionary()
	if err != nil {
		return nil, err
	}
	return &Parser{tokenizer: t, dictionary: d}, nil
}

func (p *Parser) Tokenize(text string) ([]types.Token, error) {
	tokens := p.tokenizer.Tokenize(text)
	result := make([]types.Token, 0, len(tokens))

	for _, token := range tokens {
		reading, readingExist := token.Reading()
		baseForm, baseFormExist := token.BaseForm()
		POS := token.POS()
		inflectionalForm, inflectionalFormExist := token.InflectionalForm()
		inflectionalType, inflectionalTypeExist := token.InflectionalType()
		if !readingExist {
			reading = "*"
		}
		if !baseFormExist {
			baseForm = "*"
		}
		if !inflectionalFormExist {
			inflectionalForm = "*"
		}
		if !inflectionalTypeExist {
			inflectionalType = "*"
		}

		if !kana.IsKatakana(token.Surface) {
			reading = katakanaToHiragana(reading)
		}

		result = append(result, types.Token{
			Surface: token.Surface,
			POS: POS,
			BaseForm: baseForm,
			InflectionalForm: inflectionalForm,
			InflectionalType: inflectionalType,
			Translations: p.dictionary.Lookup(token.Surface, reading),
			Reading: reading,
			Romaji: kana.KanaToRomaji(reading),
			Polivanov: "",
		})
	}
	return result, nil
}

func katakanaToHiragana(s string) string {
    runes := []rune(s)
    for i, r := range runes {
        if r >= 0x30A1 && r <= 0x30F6 {
            runes[i] = r - 0x60
        }
    }
    return string(runes)
}
