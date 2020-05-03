gen:
	protoc --go_out=plugins=grpc:internal/adapters/grpc api/*.proto

test:
	go test ./...

build: gen
	go build -o calsrv main.go
	go build -o api_server apps/api/main.go
	go build -o scheduler apps/scheduler/main.go
	go build -o sender apps/sender/main.go
