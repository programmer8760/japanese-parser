package parser

import (
	"strings"

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
		features := strings.Join(token.Features(), ",")
		result = append(result, types.Token{
			Surface: token.Surface,
			Features: features,
		})
	}
	 return result, nil
}
