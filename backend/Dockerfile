# Building stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY script.sh ./
RUN chmod +x script.sh
COPY main.sh ./
RUN chmod +x main.sh
RUN ./main.sh
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o backend .

# Deployment stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/backend .
COPY --from=builder /app/.env .
EXPOSE 1111
CMD ["./backend"]