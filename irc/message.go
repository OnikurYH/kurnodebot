package irc

import (
	"regexp"
	"strings"
)

// Message ...
type Message struct {
	Raw         string
	Prefix      string
	Nick        string
	User        string
	Host        string
	Server      string
	Command     string
	Content     string
	ContentArgs []string
	Args        []string
}

// ParseMessage ...
// Port from node-irc
func ParseMessage(rmsg string) Message {
	ircMessage := Message{
		Raw: rmsg,
	}

	reg, _ := regexp.Compile(`^:([^ ]+) +`)
	matches := reg.FindStringSubmatch(rmsg)
	// fmt.Println(matches)
	if len(matches) > 0 {
		ircMessage.Prefix = matches[1]
		reg, _ = regexp.Compile(`^:[^ ]+ +`)
		rmsg = reg.ReplaceAllString(rmsg, "")

		reg, _ = regexp.Compile("^([_a-zA-Z0-9\\~\\[\\]\\`^{}|-]*)(!([^@]+)@(.*))?$")
		matches = reg.FindStringSubmatch(ircMessage.Prefix)
		if len(matches) > 0 {
			ircMessage.Nick = matches[1]
			ircMessage.User = matches[3]
			ircMessage.Host = matches[4]
		} else {
			ircMessage.Server = ircMessage.Prefix
		}
	}

	reg, _ = regexp.Compile(`^([^ ]+) *`)
	matches = reg.FindStringSubmatch(rmsg)
	ircMessage.Command = matches[1]

	reg, _ = regexp.Compile(`^[^ ]+ +`)
	rmsg = reg.ReplaceAllString(rmsg, "")

	middle := ""
	trailing := ""
	if matched, err := regexp.MatchString(`^:|\s+:`, rmsg); err != nil || !matched {
		middle = rmsg
	} else {
		reg, _ = regexp.Compile(`(.*?)(?:^:|\s+:)(.*)`)
		matches = reg.FindStringSubmatch(rmsg)
		middle = strings.TrimRight(matches[1], " ")
		if len(matches) >= 3 {
			trailing = matches[2]
		}
	}

	ircMessage.Content = trailing
	ircMessage.ContentArgs = append([]string(nil), strings.Split(trailing, " ")...)[1:]

	if len(middle) > 0 {
		reg, _ = regexp.Compile(` +`)
		ircMessage.Args = reg.FindAllString(middle, -1)
	}

	if len(trailing) > 0 {
		ircMessage.Args = append(ircMessage.Args, trailing)
	}

	return ircMessage
}
