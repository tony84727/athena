package main

import (
	"context"
	"fmt"
	"github.com/tony84727/athena/forgegrpc"
	"google.golang.org/grpc"
	"log"
)

func main() {
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
	for {
		e, err := connect.Recv()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("sender: %s, message: %s\n", e.Sender, e.Content)
	}
}