package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	desc "github.com/xeeetu/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const address = "localhost:50051"

func main() {
	cli, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if erCli := cli.Close(); erCli != nil {
			log.Fatal(erCli)
		}
	}()

	c := desc.NewChatV1Client(cli)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := c.Create(ctx, &desc.CreateRequest{Usernames: []string{"John", "Petya", "Alex"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Create result: %v", res)

	_, err1 := c.SendMessage(ctx, &desc.SendMessageRequest{
		From:      "Alex",
		Text:      "Hello world",
		Timestamp: timestamppb.New(time.Now().UTC()),
	})

	if err1 != nil {
		log.Fatal(err1)
	}

	_, err2 := c.Delete(ctx, &desc.DeleteRequest{Id: gofakeit.Int64()})

	if err2 != nil {
		log.Fatal(err2)
	}
}
