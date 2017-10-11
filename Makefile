build:
	protoc pb/*.proto --go_out=plugins=grpc:.

server:
	go run *.go

client:
	go run client/main.go
