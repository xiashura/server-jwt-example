FROM golang:1.15

WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app cmd/service/authentication/main.go

EXPOSE 8080

CMD ["./app"]
