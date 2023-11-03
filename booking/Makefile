.PHONY: db run proto

db:
	docker-compose up -d

run:
	go run cmd/server/main.go

proto:
	protoc --go_out=. --go-grpc_out=. ./proto/booking.proto