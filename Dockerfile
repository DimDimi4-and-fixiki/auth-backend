FROM golang:1.21.8-alpine3.19 AS builder

RUN go version
RUN apk add git

COPY ./ /app
WORKDIR /app

RUN go mod download && go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./main.go
RUN chmod +x ./.bin/app

#lightweight docker container with binary
FROM alpine:3.19.0

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /app/.bin/app .
COPY --from=0 /app/.env .

EXPOSE 8080

CMD [ "./app"]
