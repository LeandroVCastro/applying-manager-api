FROM golang:1.22.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . ./

EXPOSE 8081

WORKDIR /app/internal
RUN go build -o ./internal/main.go

CMD ["ls"]
CMD ["./tmp/main"]