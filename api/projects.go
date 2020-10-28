package api

import (
	"context"
	"encoding/json"
	pb "github.com/sergeykhomenko/tm-tool/service-project/proto/project"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	address = "localhost:50051"
)

func ListProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	if r.Method == http.MethodOptions {
		return
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewProjectServiceClient(conn)

	projects, err := client.GetProjects(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalln("Project list fetch error.", err)
	}

	jsoned, err := json.Marshal(projects)

	if err != nil {
		log.Fatalln("Project list serializing error.", err)
	}

	w.Write(jsoned)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	if r.Method == http.MethodOptions {
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln("Request body reading error", err)
	}

	var project pb.ProjectEntity

	_ = json.Unmarshal(body, &project)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewProjectServiceClient(conn)

	response, _ := client.CreateProject(context.Background(), project.Proto())
	jsoned, _ := json.Marshal(response)

	w.Write(jsoned)
}
