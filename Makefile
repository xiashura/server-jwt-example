install:
	go mod download 


build:
	go build -o app cmd/service/authentication/main.go
	./app

run:
	go run app cmd/service/authentication/main.go

docker:
	docker-compose -f deployments/docker-compose/docker-compose.yml  up -d