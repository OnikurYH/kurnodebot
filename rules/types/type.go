package types

import (
	"kurnode.com/kurnodebot/config"
	"kurnode.com/kurnodebot/irc"
)

// RuleType ...
type RuleType interface {
	Do(config.RuleOutput) string
	DoWithMessage(irc.Message, config.RuleOutput) string
}
