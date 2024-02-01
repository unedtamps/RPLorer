FROM golang:1.21.6-alpine3.19 AS builder
RUN apk update
WORKDIR /app
COPY . .
RUN apk add --no-cache curl
RUN go build -o main main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate /usr/bin/migrate
COPY .env .
COPY internal/db/migration ./migration
COPY start.sh .
COPY wait-for.sh .

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "app/start.sh" ]