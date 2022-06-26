parser.proto:
	protoc -I proto --go_out=plugins=grpc:. parser.proto

crud.proto:
	protoc -I proto --go_out=plugins=grpc:. crud.proto

run:
	go run api/cmd/app/main.go
	go run parser/cmd/app/main.go
	go run crud-service/cmd/app/main.go

