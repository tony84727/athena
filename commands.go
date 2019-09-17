package main

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/gempir/go-twitch-irc"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

// CCommand stands for Calculator Command.
type CCommand struct {
}

type EmptyParameter struct {
}

func (EmptyParameter) Get(name string) (interface{}, error) {
	return nil,errors.New("unknown")
}

func (CCommand) Run(out chan<- string, args string) {
	defer close(out)
	exprStr := strings.TrimSpace(args)
	expr, err := govaluate.NewEvaluableExpression(exprStr)
	if err != nil {
		out <- "[計算模組]無法解析 " + exprStr
		log.Printf("fail to create expression %s\n", exprStr)
		log.Println(err.Error())
		return
	}
	result, err := expr.Eval(EmptyParameter{})
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
	out chan<- string
}

func (c *TwitchCommand) Run(out chan<- string, args string) {
	if c.client == nil {
		c.client = twitch.NewClient(viper.GetString("twitch.username"), viper.GetString("twitch.secret"))
		if err := c.client.Connect(); err != nil {
			log.Println("cannot connect to twitch")
			log.Println(err)
			return
		}
	}
	if c.out != nil {
		close(out)
	} else {
		c.out = out
	}
	arguments := strings.Split(args, " ")
	if len(arguments) <= 0 {
		c.out <- "錯誤的使用方式"
		return
	}
	switch arguments[0] {
	case "join":
		if len(c.channel) > 0 {
			c.client.Depart(c.channel)
		}
		channel := strings.TrimSpace(arguments[1])
		c.client.Join(channel)
		c.channel = channel
	case "out":
		if len(c.channel) <= 0 {
			return
		}
		c.client.Depart(c.channel)
		c.channel = ""
	default:
		c.out <- "錯誤的使用方式"
	}
}

func NewTwitchCommand() *TwitchCommand {
	return &TwitchCommand{
		client: twitch.NewClient(os.Getenv("TWITCH_USER"), os.Getenv("TWITCH_SECRET_KEY")),
	}
}
