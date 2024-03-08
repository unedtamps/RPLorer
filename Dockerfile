FROM golang:1.21.7-alpine3.19 AS builder
RUN apk update
WORKDIR /app
COPY . .
RUN apk add --no-cache curl
RUN go build -o main main.go
RUN go build -o seed internal/seed/main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/seed .
COPY --from=builder /app/migrate /usr/bin/migrate
COPY .env .
COPY template ./template
COPY storage ./storage
COPY internal/db/migration ./migration
COPY start.sh .
COPY wait-for.sh .

EXPOSE 8080
CMD ["/app/main" ]
ENTRYPOINT [ "app/start.sh" ]
