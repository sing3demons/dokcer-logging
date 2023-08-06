FROM golang:1.20-alpine3.16 as step_1
WORKDIR /app
COPY go.* .
RUN go mod tidy
COPY . .
RUN go build -v -o app-service

FROM alpine:3.16.0
WORKDIR /app
RUN apk add ca-certificates
COPY --from=step_1 /app/app-service .
EXPOSE 8080
CMD ["/app/app-service"]