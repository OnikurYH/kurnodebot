package main

import (
	"log"

	"kurnode.com/kurnodebot/rules"

	"kurnode.com/kurnodebot/irc"
)

// Handler ...
type Handler struct {
	client *Client
}

// NewHandler ...
func NewHandler(c *Client) *Handler {
	h := &Handler{
		client: c,
	}
	return h
}

// Listen ...
func (h *Handler) Listen() {
	go h.handle()
}

func (h *Handler) handle() {
	for {
		select {
		case data := <-h.client.OnMessage:
			im := irc.ParseMessage(data)
			h.handleMessage(im)
		}
	}
}

func (h *Handler) handleMessage(msg irc.Message) {
	log.Printf("From client [%s]", msg.Raw)
	switch msg.Command {
	case "PING":
		h.client.Write("PONG :tmi.twitch.tv")
		log.Println("PONG to server")
	case "PRIVMSG":
		h.handleUserMessage(msg)
	}
}

func (h *Handler) handleUserMessage(msg irc.Message) {
	rule, err := rules.GetFromText(msg.Content)
	if err != nil {
		return
	}
	output, err := rules.GetOutputFromIrcMessageAndRule(msg, rule.Output)
	if err != nil {
		return
	}
	h.client.SendMsg(output)
}
