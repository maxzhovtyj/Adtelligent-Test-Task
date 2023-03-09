FROM golang:1.18-alpine3.17 AS builder

RUN go version
ENV GOPATH=/

COPY . /github.com/maxzhovtyj/Adtelligent-Test-Task
WORKDIR /github.com/maxzhovtyj/Adtelligent-Test-Task

RUN go mod download
RUN GOOS=linux go build -o ./.bin/main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/maxzhovtyj/Adtelligent-Test-Task/.bin/main .
COPY --from=0 /github.com/maxzhovtyj/Adtelligent-Test-Task/configs/config.yml configs/
COPY --from=0 /github.com/maxzhovtyj/Adtelligent-Test-Task/.env .

EXPOSE 3000

CMD ["./main"]