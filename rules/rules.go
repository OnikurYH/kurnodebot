package rules

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"kurnode.com/kurnodebot/irc"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/analyzer/standard"
	"github.com/blevesearch/bleve/mapping"
	"kurnode.com/kurnodebot/config"
	"kurnode.com/kurnodebot/rules/types"
)

// RuleIndex ...
type RuleIndex struct {
	Content    string `json:"content"`
	ContentExc string `json:"content_exc"`
}

var index bleve.Index

// GetIndex ...
func GetIndex() bleve.Index {
	return index
}

// InitEngine ...
func InitEngine() error {
	RegisterSingleTokenTokenizer()
	RegisterKeywordAnalyzer()

	im := bleve.NewIndexMapping()
	im.AddDocumentMapping("doc", createDocumentMapping())
	im.DefaultType = "doc"

	idx, err := bleve.NewMemOnly(im)
	if err != nil {
		return err
	}
	index = idx

	importConfig()

	return nil
}

// GetFromText ...
func GetFromText(text string) (*config.Rule, error) {
	if strings.HasPrefix(text, "~") && strings.Contains(text, " ") {
		text = string([]rune(text)[0:strings.Index(text, " ")])
	}

	text = strings.Replace(text, " ", "\\ ", -1)
	text = strings.Replace(text, "~", "\\~", -1)

	query := bleve.NewQueryStringQuery("+content:" + text + " content_exc:" + text)
	search := bleve.NewSearchRequest(query)
	search.Size = 1
	searchResults, err := index.Search(search)
	if err != nil {
		return nil, err
	}
	if searchResults.Total < 1 {
		return nil, errors.New("No rule found")
	}
	id := searchResults.Hits[0].ID
	pos, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, err
	}
	return &config.Get().Rules[pos], nil
}

// GetRuleTypeFromConfig ...
func GetRuleTypeFromConfig(ro config.RuleOutput) (types.RuleType, error) {
	var rt types.RuleType
	switch ro.Type {
	case config.RuleOutputTypeMessage:
		rt = &types.RuleTypeMessage{}
	case config.RuleOutputTypeTime:
		rt = &types.RuleTypeTime{}

	case config.RuleOutputTypeTimeWeatherCurrent:
		rt = &types.RuleTypeCurrentWeather{}
	default:
		return nil, errors.New("No type found")
	}
	return rt, nil
}

// GetOutputFromIrcMessageAndRule ...
func GetOutputFromIrcMessageAndRule(msg irc.Message, ro config.RuleOutput) (string, error) {
	rt, err := GetRuleTypeFromConfig(ro)
	if err != nil {
		return "", err
	}
	return rt.DoWithMessage(msg, ro), nil
}

func createDocumentMapping() *mapping.DocumentMapping {
	dm := bleve.NewDocumentMapping()

	ctfm := bleve.NewTextFieldMapping()
	ctfm.Analyzer = standard.Name
	dm.AddFieldMappingsAt("content", ctfm)

	cetfm := bleve.NewTextFieldMapping()
	cetfm.Analyzer = "keyword"
	dm.AddFieldMappingsAt("content_exc", cetfm)

	return dm
}

func importConfig() {
	for i, rule := range config.Get().Rules {
		ri := RuleIndex{
			Content:    rule.Input.FormattedContent(),
			ContentExc: rule.Input.FormattedContent(),
		}
		if err := index.Index(strconv.Itoa(i), ri); err != nil {
			log.Println("[Warn] Rule cannot index " + err.Error())
		}
	}
}
