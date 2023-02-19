FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o app cmd/app/main.go

EXPOSE 8080

CMD ["./app"]
