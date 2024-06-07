FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o main ./main.go

RUN chmod +x main 

EXPOSE 8080

CMD ["./main"]