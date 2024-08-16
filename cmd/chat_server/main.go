package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	desc "github.com/xeeetu/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	address = "localhost:50051"
)

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, in *desc.CreateRequest) (*desc.CreateResponse, error) {
	select {
	case <-ctx.Done():
		// Если контекст отменён, возвращаем ошибку
		return nil, ctx.Err()
	default:
		// Продолжаем выполнение
	}

	fmt.Printf("Create chat with users: %v", in.Usernames)
	return &desc.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (s *server) Delete(ctx context.Context, in *desc.DeleteRequest) (*emptypb.Empty, error) {
	select {
	case <-ctx.Done():
		// Если контекст отменён, возвращаем ошибку
		return nil, ctx.Err()
	default:
		// Продолжаем выполнение
	}

	fmt.Printf("Delete chat with id: %d", in.Id)
	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, in *desc.SendMessageRequest) (*emptypb.Empty, error) {
	select {
	case <-ctx.Done():
		// Если контекст отменён, возвращаем ошибку
		return nil, ctx.Err()
	default:
		// Продолжаем выполнение
	}

	fmt.Printf("Send message: %v", in)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterChatV1Server(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
