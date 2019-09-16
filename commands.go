package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/gempir/go-twitch-irc"
	"log"
	"os"
	"strings"
)

// CCommand stands for Calculator Command.
type CCommand struct {
}

func (CCommand) Run(out chan<- string, args string) {
	exprStr := strings.TrimSpace(args)
	expr, err := govaluate.NewEvaluableExpression(exprStr)
	if err != nil {
		out <- "[計算模組]無法解析 " + exprStr
		log.Printf("fail to create expression %s\n", exprStr)
		log.Println(err.Error())
		return
	}
	result, err := expr.Eval(nil)
	if err != nil {
		out <- "[計算模組]無法計算 " + exprStr
		log.Printf("fail to evaluate the expression %s", exprStr)
		log.Println(err.Error())
		return
	}
	out <- fmt.Sprintf("%s => %v", exprStr, result)
}

type TwitchCommand struct {
	client *twitch.Client
	channel string
}

func (c *TwitchCommand) Run(out chan<- string, args string) {
	//if len(c.channel) > 0 {
	//	c.client.Depart(c.channel)
	//}
	//c.client.Join(strings.TrimSpace(args))
	//c.client.OnNewMessage(func(channel string, user twitch.User, message twitch.Message) {
	//	if channel == c.channel
	//})
	panic("not implemented")
}

func NewTwitchCommand() *TwitchCommand {
	return &TwitchCommand{
		client: twitch.NewClient(os.Getenv("TWITCH_USER"), os.Getenv("TWITCH_SECRET_KEY")),
	}
}
