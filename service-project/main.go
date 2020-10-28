package main

import (
	"context"
	pb "github.com/sergeykhomenko/tm-tool/service-project/proto/project"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"sync"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Project) (*pb.Project, error)
	GetAll() []*pb.Project
}

type Repository struct {
	mu       sync.RWMutex
	projects []*pb.Project
}

func (repo *Repository) Create(project *pb.Project) (*pb.Project, error) {
	dsn := "host=localhost user=postgres password=example dbname=pm_tool sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&pb.ProjectEntity{})
	db.AutoMigrate(&pb.TicketEntity{})

	storedProject, err := project.ToORM()

	if err != nil {
		log.Println("Error happen during pb to orm conversion")
		return nil, err
	}

	db.Create(&storedProject)

	//repo.mu.Lock()
	//repo.projects = append(repo.projects, project)
	//repo.mu.Unlock()

	return project, nil
}

func (repo *Repository) GetAll() []*pb.Project {
	dsn := "host=localhost user=postgres password=example dbname=pm_tool sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&pb.ProjectEntity{})

	var projects []pb.ProjectEntity
	var response []*pb.Project

	db.Find(&projects)

	for _, project := range projects {
		response = append(response, project.Proto())
	}

	return response
}

type projectService struct {
	repo repository
}

func (service *projectService) CreateProject(ctx context.Context, req *pb.Project) (*pb.Response, error) {
	project, err := service.repo.Create(req)

	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Created: true,
		Project: project,
	}, nil
}

func (service *projectService) CreateTicket(ctx context.Context, req *pb.Ticket) (*pb.Response, error) {
	dsn := "host=localhost user=postgres password=example dbname=pm_tool sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&pb.ProjectEntity{})
	db.AutoMigrate(&pb.TicketEntity{})

	var project pb.ProjectEntity

	db.Find(&project, &pb.ProjectEntity{ID: req.ProjectId})

	project.Tickets = append(project.Tickets, req.ToORM())
	db.Save(&project)

	var tickets []pb.TicketEntity
	var response []*pb.Ticket
	db.Find(&tickets)

	for _, ticket := range tickets {
		response = append(response, ticket.Proto())
	}

	return &pb.Response{Tickets: response}, nil
}

func (service *projectService) GetProjects(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	return &pb.Response{
		Projects: service.repo.GetAll(),
	}, nil
}

func (service *projectService) GetProject(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}

func (service *projectService) GetTickets(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}

func main() {
	repo := &Repository{}

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen port %v", port)
	}

	server := grpc.NewServer()

	pb.RegisterProjectServiceServer(server, &projectService{repo})
	reflection.Register(server)

	log.Println("Running on port", port)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve %v", port)
	}
}
