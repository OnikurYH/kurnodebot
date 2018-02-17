package types

import (
	"strings"

	"kurnode.com/kurnodebot/config"
	"kurnode.com/kurnodebot/irc"
)

// RuleTypeMessage ...
type RuleTypeMessage struct{}

// Do ...
func (rt *RuleTypeMessage) Do(ro config.RuleOutput) string {
	return ro.Content
}

// DoWithMessage ...
func (rt *RuleTypeMessage) DoWithMessage(msg irc.Message, ro config.RuleOutput) string {
	output := rt.Do(ro)
	output = strings.Replace(output, "{{username}}", msg.Nick, -1)
	return output
}
