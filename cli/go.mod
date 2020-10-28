module tm-tool/cli

go 1.14

require (
	github.com/sergeykhomenko/tm-tool/service-project v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.32.0
)

replace github.com/sergeykhomenko/tm-tool/service-project => ../service-project
