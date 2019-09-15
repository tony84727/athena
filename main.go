package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Knetic/govaluate"
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
	out := make(chan string, 10)
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
	for {
		e, err := connect.Recv()
		if err != nil {
			log.Fatal(err)
		}
		sender, content := decodeChatEvent(e)
		fmt.Printf("sender: %s, message: %s\n", sender, content)
		if sender == "tony84727" {
			continue
		}
		expressions := extractMathExpressions(content)
		for _, e := range expressions {
			expr, err  := govaluate.NewEvaluableExpression(e)
			if err != nil {
				log.Println(err)
				continue
			}
			result, err := expr.Eval(nil)
			if err != nil {
				log.Println(err)
				continue
			}
			out <- fmt.Sprintf("%s => %v",e, result)
		}
	}
}