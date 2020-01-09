gen:
	protoc --go_out=plugins=grpc:internal/adapters/grpc api/*.proto

test:
	go test ./...

build: gen
	go build -o calsrv main.go