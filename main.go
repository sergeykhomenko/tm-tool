package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tm-tool/api"
)

type Messsage struct {
	Author  string
	Content string
}

//type Ticket struct {
//	Author string
//	Description string
//	Assigned string
//	Messages []*Messsage
//}
//
//type Project struct {
//	Title string
//	Description string
//	Tickets []*Ticket
//}

func main() {
	var public string

	router := mux.NewRouter().StrictSlash(true)

	flag.StringVar(&public, "static files", "./client/build", "the directory to serve files from")
	flag.Parse()

	router.Use(mux.CORSMethodMiddleware(router))

	router.HandleFunc("/api/projects", api.ListProjects).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/projects", api.CreateProject).Methods(http.MethodPost)
	router.HandleFunc("/api/projects/{projectId}", api.ListProjectTickets).Methods("GET")
	router.HandleFunc("/api/projects/{projectId}", api.CreateTicket).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/tickets/{ticketId}/messages", api.ListTicketMessages).Methods("GET")
	router.HandleFunc("/api/tickets/{ticketId}", api.SubmitMessage).Methods("POST")
	router.HandleFunc("/api/tickets/{ticketId}", api.SetTicketStatus).Methods("PUT")
	router.HandleFunc("/api/tickets/{ticketId}", api.CompleteTicket).Methods("DELETE")

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(public))))

	log.Fatal(http.ListenAndServe(":8080", router))
}
