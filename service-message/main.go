package main

import (
	"context"
	pb "github.com/sergeykhomenko/tm-tool/service-message/proto/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Message) (*pb.Message, error)
	GetAll() []*pb.Message
}

type Repository struct {
	mu       sync.RWMutex
	messages []*pb.Message
}

func (repo *Repository) Create(project *pb.Message) (*pb.Message, error) {
	//repo.mu.Lock()
	//repo.projects = append(repo.projects, project)
	//repo.mu.Unlock()

	return project, nil
}

func (repo *Repository) GetAll() []*pb.Message {
	return repo.messages
}

type messageService struct {
	repo repository
}

func (service *messageService) CreateMessage(ctx context.Context, req *pb.Message) (*pb.Response, error) {
	message, err := service.repo.Create(req)

	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Created: true,
		Message: message,
	}, nil
}

func (service *messageService) GetMessages(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	return &pb.Response{
		Messages: service.repo.GetAll(),
	}, nil
}

func main() {
	repo := &Repository{}

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen port %v", port)
	}

	server := grpc.NewServer()

	pb.RegisterMessageServiceServer(server, &messageService{repo})
	reflection.Register(server)

	log.Println("Running on port", port)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve %v", port)
	}
}
