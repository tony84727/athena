package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CommanderTestSuite struct {
	suite.Suite
}

func TestCommander(t *testing.T) {
	suite.Run(t, &CommanderTestSuite{})
}

func (c *CommanderTestSuite) TestMatchMessage() {
	cases := []struct {
		input string
		command string
		args string
	}{
		{
			input: "!test",
			command: "test",
			args: "",
		},
		{
			input: "!test hello world",
			command: "test",
			args: "hello world",
		},
	}

	commander := NewCommander()
	for _, tc := range cases {
		matches := commander.matchMessage(tc.input)
		c.Require().NotNil(matches)
		c.Equal(tc.command, matches[1])
		c.Equal(tc.args, matches[2])
	}
}
