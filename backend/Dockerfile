# Building stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod backend
COPY .. .
RUN CGO_ENABLED=0 GOOS=linux go build -a cgo -o main .

# Deployment stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]