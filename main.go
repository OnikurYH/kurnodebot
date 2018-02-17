package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"kurnode.com/kurnodebot/config"
	"kurnode.com/kurnodebot/rules"
)

var c *Client

func main() {
	configPtr := flag.String("config", "config.yml", "Config path")
	flag.Parse()

	if err := config.Load(*configPtr); err != nil {
		log.Fatalln("Fail to load config -> " + err.Error())
	}

	rules.InitEngine()

	opts := ClientOptions{
		ip:       "irc.chat.twitch.tv:6667",
		username: config.Get().Server.Username,
		password: config.Get().Server.Password,
		channel:  config.Get().Server.Channel,
	}
	c = NewClient(opts)
	if err := c.Connect(); err != nil {
		log.Fatalln(err)
	}
	NewHandler(c).Listen()

	InitSchdulers(c)

	cmdHandle()
}

func cmdHandle() {
	// Command line
	reader := bufio.NewReader(os.Stdin)
	isStop := false
	for {
		if isStop {
			break
		}

		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		args := strings.Split(strings.TrimSuffix(text, "\n"), " ")
		cmd := args[0]

		switch cmd {
		case "send":
			c.SendMsg(args[1])
		case "test-rule":
			text := strings.Join(append([]string(nil), args...)[1:], " ")
			if rule, err := rules.GetFromText(text); err != nil {
				fmt.Println(err)
			} else {
				if rt, err := rules.GetRuleTypeFromConfig(rule.Output); err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Rule result found: %+v\n", rt.Do(rule.Output))
				}
			}
		case "stop":
			isStop = true
		case "help":
			showHelps()
		default:
			fmt.Printf("Command \"%s\" not found\n", cmd)
			showHelps()
		}
	}
}

func showHelps() {
	fmt.Println("send <message>      | Send message to chat room")
	fmt.Println("test-rule <message> | Test rule response")
	fmt.Println("stop                | Stop server")
	fmt.Println("help                | Show this help")
}
