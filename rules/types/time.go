package types

import (
	"strings"
	"time"

	"kurnode.com/kurnodebot/config"
	"kurnode.com/kurnodebot/irc"
)

// RuleTypeTime ...
type RuleTypeTime struct{}

// Do ...
func (rt *RuleTypeTime) Do(ro config.RuleOutput) string {
	t := time.Now().Format(ro.TimeFormat)

	output := strings.Replace(ro.Content, "{{time}}", t, -1)
	return output
}

// DoWithMessage ...
func (rt *RuleTypeTime) DoWithMessage(msg irc.Message, ro config.RuleOutput) string {
	return rt.Do(ro)
}
