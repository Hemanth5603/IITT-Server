FROM golang:latest

WORKDIR /app

ADD . .

RUN go build -o bin .

EXPOSE 80 

ENTRYPOINT [ "/app/bin" ]

