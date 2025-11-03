package parser

import (
	"japanese-parser/backend/types"

	"github.com/ikawaha/kagome-dict/ipa"
  "github.com/ikawaha/kagome/v2/tokenizer"
)

type Parser struct {
	tokenizer *tokenizer.Tokenizer
}

func NewParser() (*Parser, error) {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return nil, err
	}
	return &Parser{tokenizer: t}, nil
}

func (p *Parser) Tokenize(text string) ([]types.Token, error) {
	tokens := p.tokenizer.Tokenize(text)
	result := make([]types.Token, 0, len(tokens))

	for _, token := range tokens {
		pronunciation, pronunciationExist := token.Pronunciation()
		baseForm, baseFormExist := token.BaseForm()
		POS := token.POS()
		inflectionalForm, inflectionalFormExist := token.InflectionalForm()
		inflectionalType, inflectionalTypeExist := token.InflectionalType()
		if !pronunciationExist {
			pronunciation = "*"
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
		result = append(result, types.Token{
			Surface: token.Surface,
			Pronunciation: pronunciation,
			POS: POS,
			BaseForm: baseForm,
			InflectionalForm: inflectionalForm,
			InflectionalType: inflectionalType,
		})
	}
	return result, nil
}
