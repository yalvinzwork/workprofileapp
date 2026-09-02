FROM golang:1.27 AS builder

WORKDIR /src

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM ubuntu:26.04

WORKDIR /app

COPY --from=builder /src/app /app/app

EXPOSE 8080

CMD ["/app/app"]
