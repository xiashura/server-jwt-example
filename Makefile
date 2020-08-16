
run:
	protoc -I api/service/auth --gofast_out=plugins=grpc:pkg/api api/service/auth/api.proto
	go run cmd/service/auth/main.go


test:


build:
	protoc -I api/service/auth --gofast_out=plugins=grpc:pkg/api api/service/auth/api.proto
	go build -o app cmd/service/auth/main.go
	./app

run-docker:


build-docker:

