package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"
	pb "github.com/sergeykhomenko/tm-tool/service-project/proto/project"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "mocks/project.json"
)

func parseFile(file string) (*pb.Project, error) {
	var consignment *pb.Project
	data, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewProjectServiceClient(conn)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateProject(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	r, _ = client.GetProjects(context.Background(), &pb.GetRequest{})
	log.Println("Projects: ", r.Projects)
}
