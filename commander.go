package main

import "regexp"

var commandPattern = regexp.MustCompile(`^!(\S+)\s*(.*)`)

type Command interface {
	Run(out chan<-string, args string)
}

func NewCommander() *Commander {
	return &Commander{
		commands: make(map[string]Command),
	}
}

type Commander struct {
	commands map[string]Command
}

func (c *Commander) RegisterCommand(name string, command Command) {
	c.commands[name] = command
}

func (c Commander) matchMessage(message string) []string {
	return commandPattern.FindStringSubmatch(message)
}

func (c *Commander) Run(input <-chan string, out chan<- string) {
	for message := range input {
		matches := c.matchMessage(message)
		if matches != nil {
			cmd, exist := c.commands[matches[1]]
			if exist {
				commandOutput := make(chan string)
				go func() {
					for s := range commandOutput {
						out <- s
					}
				}()
				go cmd.Run(commandOutput, matches[2])
			}
		}
	}
}