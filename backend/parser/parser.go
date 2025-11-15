package parser

import (
	"github.com/programmer8760/japanese-parser/backend/types"
	"github.com/programmer8760/japanese-parser/backend/dictionary"
	"github.com/programmer8760/japanese-parser/backend/utils"

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

		switch token.Surface {
		case "は":
			if POS[0] == "助詞" { reading = "ワ" }
		case "へ":
				if POS[0] == "助詞" { reading = "エ" }
		case "を":
				if POS[0] == "助詞" { reading = "オ" }
		}

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
			reading = utils.KatakanaToHiragana(reading)
		}

		var translations []types.DictionaryEntry
		if token.Surface != baseForm {
			baseFormToken := p.tokenizer.Tokenize(baseForm)
			baseFormReading, baseFormReadingExists := baseFormToken[0].Reading()
			if baseFormReadingExists {
				baseFormReading = utils.KatakanaToHiragana(baseFormReading)
				translations = p.dictionary.Lookup(
					baseForm,
					baseFormReading,
				)
			} else {
				translations = p.dictionary.Lookup(
					token.Surface,
					reading,
				)
			}
		} else {
			translations = p.dictionary.Lookup(
				token.Surface,
				reading,
			)
		}

		result = append(result, types.Token{
			Surface: token.Surface,
			POS: POS,
			BaseForm: baseForm,
			InflectionalForm: inflectionalForm,
			InflectionalType: inflectionalType,
			Translations: translations,
			Reading: reading,
			Polivanov: "",
			Romaji: utils.KanaToRomaji(reading),
		})
	}
	return result, nil
}
