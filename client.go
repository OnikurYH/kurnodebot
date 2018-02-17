package main

import (
	"bufio"
	"log"
	"net"
	"net/textproto"
	"time"

	"kurnode.com/kurnodebot/config"
)

// ClientOptions ...
type ClientOptions struct {
	ip       string
	username string
	password string
	channel  string
}

// Client ...
type Client struct {
	Opts ClientOptions

	OnMessage chan string

	conn   net.Conn
	reader *textproto.Reader
	writer *bufio.Writer

	lastSendTime time.Time
}

// Connect ...
func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", c.Opts.ip)
	if err != nil {
		return err
	}

	c.conn = conn
	c.reader = textproto.NewReader(bufio.NewReader(conn))
	c.writer = bufio.NewWriter(conn)

	c.Listen()

	c.Write("PASS oauth:" + c.Opts.password)
	c.Write("NICK " + c.Opts.username)
	c.Write("JOIN #" + c.Opts.channel)

	return nil
}

// Reconnect ...
func (c *Client) Reconnect() {
	c.Disconnect()
	c.Connect()
}

// Disconnect ...
func (c *Client) Disconnect() {
	c.conn.Close()
}

// Listen ...
func (c *Client) Listen() {
	go c.listenRead()
}

// SendMsg ...
func (c *Client) SendMsg(msg string) {
	perSeconds := 30.0 / config.Get().Server.LimitSendCount
	diff := time.Since(c.lastSendTime).Seconds()
	if diff <= perSeconds {
		log.Printf("Send messsage too fast between %fs", diff)
		return
	}

	c.lastSendTime = time.Now()
	c.Write("PRIVMSG #" + c.Opts.channel + " :" + msg)
}

// Write ...
func (c *Client) Write(msg string) {
	c.writer.WriteString(msg + "\r\n")
	c.writer.Flush()
}

func (c *Client) listenRead() {
	for {
		line, err := c.reader.ReadLine()
		if err != nil {
			log.Fatalln(err)
			break
		}
		// log.Printf("%s", line)
		c.OnMessage <- line
	}
}

// NewClient ...
func NewClient(opts ClientOptions) *Client {
	c := &Client{
		Opts:      opts,
		OnMessage: make(chan string),

		lastSendTime: time.Now(),
	}
	return c
}
