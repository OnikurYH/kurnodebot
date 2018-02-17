package main

import (
	"time"

	"kurnode.com/kurnodebot/config"
	"kurnode.com/kurnodebot/rules"
)

// Scheduler ...
type Scheduler struct {
	c      *Client
	conf   config.Scheduler
	ticker *time.Ticker
}

// Start ...
func (s *Scheduler) Start() {
	s.ticker = time.NewTicker(time.Duration(s.conf.Interval) * time.Second)
	go s.tick()
}

func (s *Scheduler) tick() {
	for _ = range s.ticker.C {
		rt, err := rules.GetRuleTypeFromConfig(s.conf.Output)
		if err != nil {
			continue
		}
		str := rt.Do(s.conf.Output)
		c.SendMsg(str)
	}
}

// InitSchdulers ...
func InitSchdulers(c *Client) {
	for _, scConf := range config.Get().Schedulers {
		sc := Scheduler{
			c:    c,
			conf: scConf,
		}
		sc.Start()
	}
}
