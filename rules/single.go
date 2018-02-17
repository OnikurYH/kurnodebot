package rules

import (
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"
)

// SingleTokenTokenizer ...
type SingleTokenTokenizer struct {
}

// Tokenize ...
func (t *SingleTokenTokenizer) Tokenize(input []byte) analysis.TokenStream {
	return analysis.TokenStream{
		&analysis.Token{
			Term:     input,
			Position: 1,
			Start:    0,
			End:      len(input),
			Type:     analysis.AlphaNumeric,
		},
	}
}

// RegisterSingleTokenTokenizer ...
func RegisterSingleTokenTokenizer() {
	registry.RegisterTokenizer("single", func(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
		return &SingleTokenTokenizer{}, nil
	})
}
