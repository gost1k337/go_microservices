.PHONY: protos

run dev:
	go run cmd/app/main.go

protos:
	protoc -I protos/ --go_out=./protos --go-grpc_out=./protos protos/currency.proto