module tm-tool

go 1.14

require (
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/sergeykhomenko/tm-tool/service-project v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sys v0.0.0-20200909081042-eff7692f9009 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200911024640-645f7a48b24f // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/sergeykhomenko/tm-tool/service-project => ./service-project

replace github.com/sergeykhomenko/tm-tool/service-message => ./service-message
