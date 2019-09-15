package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/tony84727/athena/forgegrpc"
	"google.golang.org/grpc"
	"log"
	"os"
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
	//connect.Send(&forgegrpc.Message{
	//	Content:"grpc已連線",
	//})
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _, err := reader.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			connect.Send(&forgegrpc.Message{
				Content: string(input),
			})
		}

	}()
	for {
		e, err := connect.Recv()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("sender: %s, message: %s\n", e.Sender, e.Content)
	}
}