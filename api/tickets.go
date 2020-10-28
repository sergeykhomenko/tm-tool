package api

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/sergeykhomenko/tm-tool/service-project/proto/project"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	projectService = "localhost:50051"
)

func ListProjectTickets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	if r.Method == http.MethodOptions {
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln("Request body reading error", err)
	}

	var ticket pb.TicketEntity

	_ = json.Unmarshal(body, &ticket)

	conn, err := grpc.Dial(projectService, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewProjectServiceClient(conn)

	response, _ := client.CreateTicket(context.Background(), ticket.Proto())
	jsoned, _ := json.Marshal(response)

	w.Write(jsoned)
}

func SetTicketStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func CompleteTicket(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
