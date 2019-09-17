package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tony84727/athena/forgegrpc"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	viper.SetConfigName("athena")
	viper.AddConfigPath(".")
	viper.AddConfigPath(os.Getenv("BUILD_WORKSPACE_DIRECTORY"))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Println("error on load config")
			log.Println(err)
		}
	}
	client, err := grpc.Dial("localhost:30000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	chat := forgegrpc.NewChatClient(client)
	ctx := context.Background()
	connect, err := chat.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	out := make(chan string, 100)
	defer close(out)
	go func() {
		for msg := range out {
			connect.Send(&forgegrpc.Message{
				Content: msg,
			})
		}
	}()
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _, err := reader.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			out <- string(input)
		}
	}()
	commander := NewCommander()
	commander.RegisterCommand("c", CCommand{})
	commander.RegisterCommand("twitch", NewTwitchCommand())
	commanderInput := make(chan string, 100)
	go commander.Run(commanderInput, out)
	for {
		e, err := connect.Recv()
		if err != nil {
			log.Fatal(err)
		}
		sender, content := decodeChatEvent(e)
		fmt.Printf("sender: %s, message: %s\n", sender, content)
		commanderInput <- content
	}
}