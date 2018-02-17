package rules

import (
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"
)

// RegisterKeywordAnalyzer ...
func RegisterKeywordAnalyzer() {
	registry.RegisterAnalyzer("keyword", func(config map[string]interface{}, cache *registry.Cache) (*analysis.Analyzer, error) {
		keywordTokenizer, err := cache.TokenizerNamed("single")
		if err != nil {
			return nil, err
		}
		rv := analysis.Analyzer{
			Tokenizer: keywordTokenizer,
		}
		return &rv, nil
	})
}
