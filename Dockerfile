FROM golang:latest

WORKDIR /app

COPY go.mod go.sum go.work ./

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

