FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o webapp cmd/webapp/main.go

EXPOSE 8080

CMD ["./webapp"]
