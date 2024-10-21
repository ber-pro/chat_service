package main

import (
	desc "chat_server/pkg/chat_v1"
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type Server struct {
	desc.UnimplementedNoteV1Server
}

const grpcPort = 8800

func (s *Server) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}
func (s *Server) Delete(ctx context.Context, request *desc.DeleteRequest) (*emptypb.Empty, error) {
	var empty *emptypb.Empty
	log.Printf("Message with id=%d", request.Id)
	return empty, nil
}
func (s *Server) SendMessage(ctx context.Context, request *desc.SendMessageRequest) (*emptypb.Empty, error) {
	var empty *emptypb.Empty
	log.Printf("Message from %s send", request.From)
	return empty, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failt to listen: %v")
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &Server{})

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
