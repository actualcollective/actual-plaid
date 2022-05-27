FROM golang:1.18-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o ./main internal/main.go

CMD ["/app/main"]