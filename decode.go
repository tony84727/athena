package main

import (
	"github.com/tony84727/athena/forgegrpc"
	"regexp"
)

var messageDecodePattern = regexp.MustCompile(`<§\w(.*)§\w>§\w\s*§\w(.*)§\w`)

func decodeChatEvent(e *forgegrpc.ChatEvent) (sender string, content string) {
	if len(e.Sender) > 0 {
		return e.Sender, e.Content
	}
	matches := messageDecodePattern.FindStringSubmatch(e.Content)
	if matches == nil {
		return "",""
	}
	return matches[1], matches[2]
}
